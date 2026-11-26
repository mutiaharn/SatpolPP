package handlers

import (
	"fmt"
	"net/http"

	"satpolpp-sulsel-api/internal/adapters/storage"
	"satpolpp-sulsel-api/internal/contracts"
	"satpolpp-sulsel-api/internal/utils"
)

type UploadHandler struct {
	storage *storage.UnifiedStorage
}

func NewUploadHandler(st *storage.UnifiedStorage) *UploadHandler {
	return &UploadHandler{storage: st}
}

func (h *UploadHandler) UploadEvidence(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, contracts.Shared.UploadLimits.MaxImageSizeBytes+1024*1024)
	if err := r.ParseMultipartForm(contracts.Shared.UploadLimits.MaxImageSizeBytes); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "FILE_TOO_LARGE", "Ukuran file foto melebihi batas maksimal 5 MB")
		return
	}

	_, fh, err := r.FormFile("foto")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "MISSING_FILE", "Berkas foto wajib diunggah pada field 'foto'")
		return
	}

	url, err := h.storage.UploadEvidenceImage(r.Context(), fh)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_IMAGE", err.Error())
		return
	}

	utils.RespondCreated(w, map[string]any{
		"success":   true,
		"image_url": url,
	})
}

func (h *UploadHandler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, contracts.Shared.UploadLimits.MaxImageSizeBytes+1024*1024)
	if err := r.ParseMultipartForm(contracts.Shared.UploadLimits.MaxImageSizeBytes); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "FILE_TOO_LARGE", "Ukuran gambar melebihi batas maksimal 5 MB")
		return
	}

	_, fh, err := r.FormFile("file")
	if err != nil {
		_, fh, err = r.FormFile("foto")
	}
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "MISSING_FILE", "Berkas gambar wajib diunggah")
		return
	}

	url, err := h.storage.UploadMediaImage(r.Context(), fh)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_IMAGE", err.Error())
		return
	}

	utils.RespondCreated(w, map[string]any{
		"success": true,
		"url":     url,
	})
}

func (h *UploadHandler) UploadDocument(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, contracts.Shared.UploadLimits.MaxDocSizeBytes+1024*1024)
	if err := r.ParseMultipartForm(contracts.Shared.UploadLimits.MaxDocSizeBytes); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "FILE_TOO_LARGE", "Ukuran dokumen melebihi batas maksimal 20 MB")
		return
	}

	_, fh, err := r.FormFile("file")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "MISSING_FILE", "Berkas dokumen wajib diunggah pada field 'file'")
		return
	}

	url, sizeBytes, err := h.storage.UploadPublicDocument(r.Context(), fh)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_DOCUMENT", err.Error())
		return
	}

	sizeLabel := fmt.Sprintf("%.1f KB", float64(sizeBytes)/1024.0)
	if sizeBytes >= 1024*1024 {
		sizeLabel = fmt.Sprintf("%.2f MB", float64(sizeBytes)/(1024.0*1024.0))
	}

	utils.RespondCreated(w, map[string]any{
		"success":   true,
		"file_url":  url,
		"file_size": sizeLabel,
	})
}
