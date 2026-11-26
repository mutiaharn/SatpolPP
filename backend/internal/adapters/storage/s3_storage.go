package storage

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"satpolpp-sulsel-api/internal/config"
	"satpolpp-sulsel-api/internal/contracts"
)

var (
	ErrInvalidDocType = errors.New("berkas harus berupa PDF, Word (.doc/.docx), atau Excel (.xls/.xlsx)")
)

var docMIMEExtensions = map[string]string{
	"application/pdf": ".pdf",
	"application/msword": ".doc",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",
	"application/vnd.ms-excel": ".xls",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": ".xlsx",
}

// UnifiedStorage abstracts RustFS S3-compatible storage with local disk fallback for dev.
type UnifiedStorage struct {
	s3          *minio.Client
	localDir    string
	evidenceBkt string
	docsBkt     string
	mediaBkt    string
}

func NewUnifiedStorage(cfg *config.Config) (*UnifiedStorage, error) {
	u := &UnifiedStorage{
		localDir:    cfg.StoragePath,
		evidenceBkt: cfg.S3BucketEvidence,
		docsBkt:     cfg.S3BucketDocs,
		mediaBkt:    cfg.S3BucketMedia,
	}

	if cfg.S3Endpoint != "" && cfg.S3AccessKey != "" {
		endpoint := strings.TrimPrefix(strings.TrimPrefix(cfg.S3Endpoint, "https://"), "http://")
		client, err := minio.New(endpoint, &minio.Options{
			Creds:        credentials.NewStaticV4(cfg.S3AccessKey, cfg.S3SecretKey, ""),
			Secure:       cfg.S3UseSSL,
			BucketLookup: minio.BucketLookupPath, // Required for RustFS / MinIO path-style
		})
		if err != nil {
			return nil, fmt.Errorf("failed to initialize RustFS S3 client: %w", err)
		}
		u.s3 = client
		if err := u.ensureBuckets(context.Background()); err != nil {
			return nil, fmt.Errorf("failed to ensure RustFS buckets: %w", err)
		}
	} else {
		if err := os.MkdirAll(cfg.StoragePath, 0755); err != nil {
			return nil, fmt.Errorf("failed to create local storage dir: %w", err)
		}
	}

	return u, nil
}

// UploadEvidenceImage stores private citizen complaint photos (max 5MB, JPEG/PNG/WebP).
func (u *UnifiedStorage) UploadEvidenceImage(ctx context.Context, fh *multipart.FileHeader) (string, error) {
	return u.uploadValidated(ctx, u.evidenceBkt, "evidence", fh, contracts.Shared.UploadLimits.MaxImageSizeBytes, allowedMIMEs, false)
}

// UploadMediaImage stores public portal images (thumbnails, galleries, leadership).
func (u *UnifiedStorage) UploadMediaImage(ctx context.Context, fh *multipart.FileHeader) (string, error) {
	return u.uploadValidated(ctx, u.mediaBkt, "media", fh, contracts.Shared.UploadLimits.MaxImageSizeBytes, allowedMIMEs, true)
}

// UploadPublicDocument stores official government documents (PDF, DOC/DOCX, XLS/XLSX).
func (u *UnifiedStorage) UploadPublicDocument(ctx context.Context, fh *multipart.FileHeader) (string, int64, error) {
	key, err := u.uploadValidated(ctx, u.docsBkt, "documents", fh, contracts.Shared.UploadLimits.MaxDocSizeBytes, docMIMEExtensions, true)
	return key, fh.Size, err
}

func (u *UnifiedStorage) uploadValidated(
	ctx context.Context,
	bucket, subFolder string,
	fh *multipart.FileHeader,
	maxBytes int64,
	allowed map[string]string,
	isPublic bool,
) (string, error) {
	if fh.Size == 0 {
		return "", ErrEmptyFile
	}
	if fh.Size > maxBytes {
		return "", fmt.Errorf("ukuran berkas melebihi batas %d MB", maxBytes/(1024*1024))
	}

	f, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open upload: %w", err)
	}
	defer f.Close()

	head := make([]byte, 512)
	n, err := f.Read(head)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to inspect file header: %w", err)
	}

	detected := http.DetectContentType(head[:n])
	if idx := strings.Index(detected, ";"); idx != -1 {
		detected = detected[:idx]
	}

	// Office OpenXML (.docx/.xlsx) are ZIP containers detected as application/zip by net/http
	if detected == "application/zip" {
		lowerExt := strings.ToLower(filepath.Ext(fh.Filename))
		if lowerExt == ".docx" {
			detected = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
		} else if lowerExt == ".xlsx" {
			detected = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
		}
	}

	ext, ok := allowed[detected]
	if !ok {
		return "", fmt.Errorf("tipe berkas tidak diizinkan (%s)", detected)
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	rnd := make([]byte, 16)
	_, _ = rand.Read(rnd)
	objectKey := fmt.Sprintf("%s/%d_%s%s", time.Now().Format("2006/01"), time.Now().Unix(), hex.EncodeToString(rnd), ext)

	if u.s3 != nil {
		_, err := u.s3.PutObject(ctx, bucket, objectKey, f, fh.Size, minio.PutObjectOptions{
			ContentType: detected,
		})
		if err != nil {
			return "", fmt.Errorf("rustfs upload failed: %w", err)
		}
		if isPublic {
			return fmt.Sprintf("/storage/%s/%s", bucket, objectKey), nil
		}
		return fmt.Sprintf("/api/v1/admin/evidence/%s", objectKey), nil
	}

	// Local disk fallback (dev mode without Docker/RustFS)
	diskSub := filepath.Join(u.localDir, subFolder)
	_ = os.MkdirAll(diskSub, 0755)
	flatName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), hex.EncodeToString(rnd), ext)
	dstPath := filepath.Join(diskSub, flatName)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, f); err != nil {
		return "", err
	}

	if isPublic {
		return fmt.Sprintf("/api/v1/uploads/%s/%s", subFolder, flatName), nil
	}
	return fmt.Sprintf("/api/v1/admin/evidence/local/%s", flatName), nil
}

// GetPrivateEvidence streams an image from the private RustFS bucket or local fallback.
func (u *UnifiedStorage) GetPrivateEvidence(ctx context.Context, key string) (io.ReadCloser, string, error) {
	cleanKey := strings.TrimPrefix(filepath.ToSlash(filepath.Clean("/"+key)), "/")

	if u.s3 != nil && !strings.HasPrefix(cleanKey, "local/") {
		obj, err := u.s3.GetObject(ctx, u.evidenceBkt, cleanKey, minio.GetObjectOptions{})
		if err != nil {
			return nil, "", err
		}
		info, err := obj.Stat()
		if err != nil {
			_ = obj.Close()
			return nil, "", err
		}
		return obj, info.ContentType, nil
	}

	localName := filepath.Base(strings.TrimPrefix(cleanKey, "local/"))
	fullPath := filepath.Join(u.localDir, "evidence", localName)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, "", err
	}
	mimeType := http.DetectContentType(data)
	return io.NopCloser(bytes.NewReader(data)), mimeType, nil
}

func (u *UnifiedStorage) ensureBuckets(ctx context.Context) error {
	buckets := []struct {
		name     string
		isPublic bool
	}{
		{u.evidenceBkt, false},
		{u.docsBkt, true},
		{u.mediaBkt, true},
	}

	for _, b := range buckets {
		if b.name == "" {
			continue
		}
		exists, err := u.s3.BucketExists(ctx, b.name)
		if err != nil {
			return err
		}
		if !exists {
			if err := u.s3.MakeBucket(ctx, b.name, minio.MakeBucketOptions{}); err != nil {
				return err
			}
		}
		if b.isPublic {
			policy := fmt.Sprintf(`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::%s/*"]}]}`, b.name)
			_ = u.s3.SetBucketPolicy(ctx, b.name, policy)
		}
	}
	return nil
}
