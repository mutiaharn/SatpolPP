package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"satpolpp-sulsel-api/internal/adapters/middleware"
	"satpolpp-sulsel-api/internal/adapters/repository/postgres"
	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

type PortalHandler struct {
	db *postgres.DB
}

func NewPortalHandler(db *postgres.DB) *PortalHandler {
	return &PortalHandler{db: db}
}

func (h *PortalHandler) ListRegencies(w http.ResponseWriter, r *http.Request) {
	list, err := h.db.ListRegencies(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat daftar wilayah")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": list})
}

func (h *PortalHandler) GetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.db.GetSettingsMap(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat konfigurasi portal")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "settings": settings})
}

func (h *PortalHandler) UpdateSettings(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Payload pengaturan tidak valid")
		return
	}
	if err := h.db.UpsertSettings(r.Context(), req); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menyimpan pengaturan portal")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "settings": req})
}

func (h *PortalHandler) ListArticles(w http.ResponseWriter, r *http.Request) {
	list, err := h.db.ListArticles(r.Context(), true)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat artikel")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": list})
}

func (h *PortalHandler) GetArticleBySlug(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimSpace(r.PathValue("slug"))
	art, err := h.db.GetArticleBySlug(r.Context(), slug)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "NOT_FOUND", "Artikel tidak ditemukan")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat detail artikel")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": art})
}

func (h *PortalHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var a domain.Article
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Payload artikel tidak valid")
		return
	}
	if strings.TrimSpace(a.Title) == "" || strings.TrimSpace(a.Content) == "" {
		utils.RespondError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Judul dan konten artikel wajib diisi")
		return
	}
	if strings.TrimSpace(a.Slug) == "" {
		a.Slug = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(a.Title), " ", "-"))
	}
	if actor, ok := middleware.GetAuthUser(r.Context()); ok {
		a.AuthorID = actor.ID
		a.AuthorName = actor.Name
	} else {
		a.AuthorID = 1
	}
	a.IsPublished = true

	if err := h.db.CreateArticle(r.Context(), &a); err != nil {
		utils.RespondError(w, http.StatusConflict, "CREATE_FAILED", "Gagal menyimpan artikel (slug mungkin sudah digunakan)")
		return
	}
	utils.RespondCreated(w, map[string]any{"success": true, "data": a})
}

func (h *PortalHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_ID", "ID artikel tidak valid")
		return
	}
	if err := h.db.DeleteArticle(r.Context(), uint(id)); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "NOT_FOUND", "Artikel tidak ditemukan")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menghapus artikel")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true})
}

func (h *PortalHandler) ListDocuments(w http.ResponseWriter, r *http.Request) {
	cat := strings.TrimSpace(r.URL.Query().Get("category"))
	list, err := h.db.ListDocuments(r.Context(), cat)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat dokumen publik")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": list})
}

func (h *PortalHandler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	var doc domain.PublicDocument
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Payload dokumen tidak valid")
		return
	}
	if strings.TrimSpace(doc.Title) == "" || strings.TrimSpace(doc.FileURL) == "" {
		utils.RespondError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Judul dan URL berkas dokumen wajib diisi")
		return
	}
	if doc.Year == 0 {
		doc.Year = time.Now().Year()
	}
	if err := h.db.CreateDocument(r.Context(), &doc); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menyimpan dokumen publik")
		return
	}
	utils.RespondCreated(w, map[string]any{"success": true, "data": doc})
}

func (h *PortalHandler) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_ID", "ID dokumen tidak valid")
		return
	}
	if err := h.db.DeleteDocument(r.Context(), uint(id)); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menghapus dokumen")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true})
}

func (h *PortalHandler) ListGalleries(w http.ResponseWriter, r *http.Request) {
	list, err := h.db.ListGalleries(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat galeri")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": list})
}

func (h *PortalHandler) CreateGallery(w http.ResponseWriter, r *http.Request) {
	var g domain.Gallery
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Payload galeri tidak valid")
		return
	}
	if g.EventDate.IsZero() {
		g.EventDate = time.Now()
	}
	if err := h.db.CreateGallery(r.Context(), &g); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menyimpan item galeri")
		return
	}
	utils.RespondCreated(w, map[string]any{"success": true, "data": g})
}

func (h *PortalHandler) DeleteGallery(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_ID", "ID galeri tidak valid")
		return
	}
	if err := h.db.DeleteGallery(r.Context(), uint(id)); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menghapus galeri")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true})
}

func (h *PortalHandler) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	list, err := h.db.ListOrganizations(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat struktur SOTK")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": list})
}

func (h *PortalHandler) UpsertOrganization(w http.ResponseWriter, r *http.Request) {
	var node domain.OrganizationNode
	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Payload SOTK tidak valid")
		return
	}
	if strings.TrimSpace(node.OrgCode) == "" || strings.TrimSpace(node.PositionTitle) == "" {
		utils.RespondError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Kode unit dan nama jabatan wajib diisi")
		return
	}
	if node.SubUnits == nil {
		node.SubUnits = []string{}
	}
	if err := h.db.UpsertOrganization(r.Context(), &node); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menyimpan node SOTK")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": node})
}

func (h *PortalHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_ID", "ID unit SOTK tidak valid")
		return
	}
	if err := h.db.DeleteOrganization(r.Context(), uint(id)); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal menghapus unit SOTK")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true})
}
