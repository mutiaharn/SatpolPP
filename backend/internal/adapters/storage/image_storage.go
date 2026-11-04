package storage

import (
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
)

const (
	// MaxImageSize 5 Megabytes limit as per SIAP TERTIB requirements
	MaxImageSize = 5 * 1024 * 1024
)

var (
	ErrFileTooLarge     = errors.New("ukuran berkas melebihi batas maksimum 5 MB")
	ErrInvalidImageType = errors.New("berkas harus berupa format gambar (JPEG, PNG, atau WebP)")
	ErrEmptyFile        = errors.New("berkas foto kosong")
)

// Allowed MIME types strictly for image evidence
var allowedMIMEs = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

// ImageStorage handles safe uploading and validation of photo evidence
type ImageStorage struct {
	baseDir string
}

func NewImageStorage(baseDir string) (*ImageStorage, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload dir: %w", err)
	}
	return &ImageStorage{baseDir: baseDir}, nil
}

// SaveUploadedImage validates MIME magic bytes, size, and saves to disk securely
func (s *ImageStorage) SaveUploadedImage(fileHeader *multipart.FileHeader) (string, error) {
	if fileHeader.Size > MaxImageSize {
		return "", ErrFileTooLarge
	}
	if fileHeader.Size == 0 {
		return "", ErrEmptyFile
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer file.Close()

	// Read first 512 bytes for MIME Magic Byte Detection
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read file header: %w", err)
	}

	detectedMIME := http.DetectContentType(buffer[:n])
	// Strip parameters like charset
	if idx := strings.Index(detectedMIME, ";"); idx != -1 {
		detectedMIME = detectedMIME[:idx]
	}

	ext, isAllowed := allowedMIMEs[detectedMIME]
	if !isAllowed {
		return "", fmt.Errorf("%w: terdeteksi tipe %s", ErrInvalidImageType, detectedMIME)
	}

	// Rewind file reader back to start
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to rewind file: %w", err)
	}

	// Generate cryptographically random unguessable filename
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random token: %w", err)
	}
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), hex.EncodeToString(randomBytes), ext)
	dstPath := filepath.Join(s.baseDir, filename)

	// Save to disk
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to write file to disk: %w", err)
	}

	// Return public accessible URL path
	publicURL := filepath.ToSlash(filepath.Join("/uploads/evidence", filename))
	return publicURL, nil
}
