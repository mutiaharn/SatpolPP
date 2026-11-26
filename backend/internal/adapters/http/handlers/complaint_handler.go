package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"satpolpp-sulsel-api/internal/adapters/middleware"
	"satpolpp-sulsel-api/internal/adapters/repository/postgres"
	"satpolpp-sulsel-api/internal/adapters/storage"
	"satpolpp-sulsel-api/internal/config"
	"satpolpp-sulsel-api/internal/contracts"
	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

type ComplaintHandler struct {
	db      *postgres.DB
	storage *storage.UnifiedStorage
	cfg     *config.Config
}

func NewComplaintHandler(db *postgres.DB, st *storage.UnifiedStorage, cfg *config.Config) *ComplaintHandler {
	return &ComplaintHandler{db: db, storage: st, cfg: cfg}
}

func isValidRegency(id string) (string, bool) {
	for _, r := range contracts.Shared.Regencies {
		if r.ID == id {
			return r.Name, true
		}
	}
	return "", false
}

func maskRecord(rec *postgres.ComplaintRecord, hideAllAnon bool) *postgres.ComplaintRecord {
	copyRec := *rec
	if rec.IsAnonymous || hideAllAnon {
		copyRec.ReporterName = utils.MaskName(rec.ReporterName)
		copyRec.ReporterNIK = rec.ReporterNIKMask
		copyRec.ReporterPhone = utils.MaskPhone("")
		copyRec.ReporterEmail = utils.MaskEmail(rec.ReporterEmail)
	} else {
		copyRec.ReporterNIK = rec.ReporterNIKMask
	}
	return &copyRec
}

func (h *ComplaintHandler) SubmitComplaint(w http.ResponseWriter, r *http.Request) {
	var req struct {
		IsAnonymous      bool     `json:"is_anonymous"`
		ReporterNIK      string   `json:"reporter_nik"`
		ReporterName     string   `json:"reporter_name"`
		ReporterPhone    string   `json:"reporter_phone"`
		ReporterEmail    string   `json:"reporter_email"`
		ViolationType    string   `json:"violation_type"`
		RegencyID        string   `json:"regency_id"`
		DistrictName     string   `json:"district_name"`
		VillageName      string   `json:"village_name"`
		AddressDetail    string   `json:"address_detail"`
		Latitude         float64  `json:"latitude"`
		Longitude        float64  `json:"longitude"`
		Description      string   `json:"description"`
		EvidenceImageURL string   `json:"evidence_image_url"`
		EvidencePhotos   []string `json:"evidence_photos"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Format payload pengaduan tidak valid")
		return
	}

	if strings.TrimSpace(req.Description) == "" || strings.TrimSpace(req.AddressDetail) == "" || strings.TrimSpace(req.ViolationType) == "" {
		utils.RespondError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Jenis pelanggaran, alamat detail, dan kronologi wajib diisi")
		return
	}

	regName, validReg := isValidRegency(req.RegencyID)
	if !validReg {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_REGENCY", "Kode kabupaten/kota Sulawesi Selatan tidak valid")
		return
	}

	nikClean := strings.TrimSpace(req.ReporterNIK)
	if nikClean != "" && len(nikClean) != 16 {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_NIK", "NIK harus terdiri dari tepat 16 digit angka")
		return
	}

	nikEnc, err := utils.EncryptPII(nikClean, h.cfg.PIIEncryptionKey)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "CRYPTO_ERROR", "Gagal mengamankan data identitas pelapor")
		return
	}
	phoneEnc, err := utils.EncryptPII(strings.TrimSpace(req.ReporterPhone), h.cfg.PIIEncryptionKey)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "CRYPTO_ERROR", "Gagal mengamankan nomor telepon pelapor")
		return
	}

	photos := req.EvidencePhotos
	if len(photos) == 0 && strings.TrimSpace(req.EvidenceImageURL) != "" {
		photos = []string{strings.TrimSpace(req.EvidenceImageURL)}
	}

	ticketCode := utils.GenerateTicketCode()
	rec := &postgres.ComplaintRecord{
		Complaint: domain.Complaint{
			TicketCode:       ticketCode,
			IsAnonymous:      req.IsAnonymous,
			ReporterName:     strings.TrimSpace(req.ReporterName),
			ReporterEmail:    strings.TrimSpace(req.ReporterEmail),
			ViolationType:    strings.TrimSpace(req.ViolationType),
			RegencyID:        req.RegencyID,
			RegencyName:      regName,
			DistrictName:     strings.TrimSpace(req.DistrictName),
			VillageName:      strings.TrimSpace(req.VillageName),
			AddressDetail:    strings.TrimSpace(req.AddressDetail),
			Latitude:         req.Latitude,
			Longitude:        req.Longitude,
			Description:      strings.TrimSpace(req.Description),
			IncidentDate:     time.Now(),
			EvidenceImageURL: "",
			Status:           domain.StatusPending,
		},
		ReporterNIKEnc:   nikEnc,
		ReporterPhoneEnc: phoneEnc,
		ReporterNIKMask:  utils.MaskNIK(nikClean),
	}
	if len(photos) > 0 {
		rec.EvidenceImageURL = photos[0]
	}

	subject, bodyText, bodyHTML := utils.BuildTicketReceiptEmail(
		h.cfg.PublicSiteURL, ticketCode, string(domain.StatusPending), regName, "", "",
	)

	if err := h.db.CreateComplaintTx(r.Context(), rec, photos, subject, bodyText, bodyHTML); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menyimpan pengaduan ke database")
		return
	}

	emailNotice := utils.EmailNotice{
		Sent:    rec.ReporterEmail != "",
		Mode:    "queued",
		To:      rec.ReporterEmail,
		Subject: subject,
		Message: "Resi tiket dijadwalkan untuk pengiriman otomatis melalui worker.",
	}
	if rec.ReporterEmail == "" {
		emailNotice.Mode = "skipped"
		emailNotice.Message = "Pelapor tidak mencantumkan alamat email."
	}

	utils.RespondCreated(w, map[string]any{
		"success":     true,
		"ticket_code": rec.TicketCode,
		"message":     "Laporan berhasil dicatat dalam sistem SIAP TERTIB",
		"email":       emailNotice,
		"data":        maskRecord(rec, true),
	})
}

func (h *ComplaintHandler) TrackComplaint(w http.ResponseWriter, r *http.Request) {
	code := strings.ToUpper(strings.TrimSpace(r.PathValue("code")))
	if code == "" {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_TICKET", "Kode tiket wajib diisi")
		return
	}

	rec, err := h.db.GetComplaintByTicket(r.Context(), code)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "TICKET_NOT_FOUND", "Kode tiket tidak ditemukan")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal melacak pengaduan")
		return
	}

	utils.RespondOK(w, map[string]any{
		"success": true,
		"data":    maskRecord(rec, true),
	})
}

func (h *ComplaintHandler) ListAdminComplaints(w http.ResponseWriter, r *http.Request) {
	statusFilter := strings.TrimSpace(r.URL.Query().Get("status"))
	list, err := h.db.ListComplaints(r.Context(), statusFilter)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat daftar pengaduan")
		return
	}

	masked := make([]*postgres.ComplaintRecord, 0, len(list))
	for i := range list {
		masked = append(masked, maskRecord(&list[i], true))
	}

	utils.RespondOK(w, map[string]any{
		"success": true,
		"data":    masked,
	})
}

func (h *ComplaintHandler) UnmaskComplaintIdentity(w http.ResponseWriter, r *http.Request) {
	ticket := strings.ToUpper(strings.TrimSpace(r.PathValue("ticket")))
	var req struct {
		Reason string `json:"reason"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if strings.TrimSpace(req.Reason) == "" {
		req.Reason = "Verifikasi lapangan / tindak lanjut resmi petugas"
	}

	rec, err := h.db.GetComplaintByTicket(r.Context(), ticket)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "TICKET_NOT_FOUND", "Tiket tidak ditemukan")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal mengambil data pengaduan")
		return
	}

	nikPlain, err := utils.DecryptPII(rec.ReporterNIKEnc, h.cfg.PIIEncryptionKey)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DECRYPT_ERROR", "Gagal mendekripsi NIK pelapor")
		return
	}
	phonePlain, err := utils.DecryptPII(rec.ReporterPhoneEnc, h.cfg.PIIEncryptionKey)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DECRYPT_ERROR", "Gagal mendekripsi nomor telepon pelapor")
		return
	}

	if actor, ok := middleware.GetAuthUser(r.Context()); ok {
		_ = h.db.RecordAuditLog(r.Context(), &domain.AuditLog{
			UserID:    actor.ID,
			UserName:  actor.Username,
			Action:    "UNMASK_PII",
			Resource:  ticket,
			Details:   fmt.Sprintf("Membuka identitas pelapor. Alasan: %s", req.Reason),
			IPAddress: r.RemoteAddr,
		})
	}

	utils.RespondOK(w, map[string]any{
		"success": true,
		"data": map[string]any{
			"ticket_code":    rec.TicketCode,
			"reporter_name":  rec.ReporterName,
			"reporter_nik":   nikPlain,
			"reporter_phone": phonePlain,
			"reporter_email": rec.ReporterEmail,
		},
	})
}

func (h *ComplaintHandler) UpdateComplaintStatus(w http.ResponseWriter, r *http.Request) {
	ticket := strings.ToUpper(strings.TrimSpace(r.PathValue("ticket")))
	var req struct {
		Status      domain.ComplaintStatus `json:"status"`
		DisposedTo  string                 `json:"disposed_to"`
		ActionNotes string                 `json:"action_notes"`
		ActionProof string                 `json:"action_proof"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Format pembaruan status tidak valid")
		return
	}

	if _, ok := contracts.Shared.ComplaintStatuses[string(req.Status)]; !ok {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_STATUS", "Status pengaduan tidak dikenali")
		return
	}

	actor, _ := middleware.GetAuthUser(r.Context())
	subject, bodyText, bodyHTML := utils.BuildTicketReceiptEmail(
		h.cfg.PublicSiteURL, ticket, string(req.Status), "", req.DisposedTo, req.ActionNotes,
	)

	updated, err := h.db.UpdateComplaintStatusTx(
		r.Context(), ticket, req.Status, req.DisposedTo, req.ActionNotes, req.ActionProof,
		actor.ID, actor.Username, actor.Name, r.RemoteAddr,
		subject, bodyText, bodyHTML,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "TICKET_NOT_FOUND", "Tiket tidak ditemukan")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memperbarui status pengaduan")
		return
	}

	emailNotice := utils.EmailNotice{
		Sent:    updated.ReporterEmail != "",
		Mode:    "queued",
		To:      updated.ReporterEmail,
		Subject: subject,
		Message: "Pembaruan status dijadwalkan untuk pengiriman email.",
	}
	if updated.ReporterEmail == "" {
		emailNotice.Mode = "skipped"
	}

	utils.RespondOK(w, map[string]any{
		"success": true,
		"email":   emailNotice,
		"data":    maskRecord(updated, true),
	})
}

func (h *ComplaintHandler) StreamPrivateEvidence(w http.ResponseWriter, r *http.Request) {
	objectKey := strings.TrimPrefix(r.URL.Path, "/api/v1/admin/evidence/")
	if objectKey == "" {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_KEY", "Object key bukti wajib disertakan")
		return
	}

	reader, mimeType, err := h.storage.GetPrivateEvidence(r.Context(), objectKey)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "NOT_FOUND", "Berkas bukti tidak ditemukan")
		return
	}
	defer reader.Close()

	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}
	w.Header().Set("Cache-Control", "private, max-age=300")
	_, _ = io.Copy(w, reader)
}
