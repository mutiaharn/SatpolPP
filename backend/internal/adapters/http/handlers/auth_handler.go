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
	"satpolpp-sulsel-api/internal/config"
	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

type AuthHandler struct {
	db  *postgres.DB
	cfg *config.Config
}

func NewAuthHandler(db *postgres.DB, cfg *config.Config) *AuthHandler {
	return &AuthHandler{db: db, cfg: cfg}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Payload login tidak valid")
		return
	}

	user, err := h.db.GetUserByUsername(r.Context(), strings.TrimSpace(req.Username))
	if err != nil || !utils.VerifyPassword(req.Password, user.Password) {
		utils.RespondError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Username atau password salah")
		return
	}
	if !user.IsActive {
		utils.RespondError(w, http.StatusForbidden, "ACCOUNT_DISABLED", "Akun ini telah dinonaktifkan oleh Superadmin")
		return
	}

	token, exp, err := utils.GenerateJWT(user, h.cfg.JWTSecret, 8*time.Hour)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "TOKEN_ERROR", "Gagal menerbitkan sesi autentikasi")
		return
	}

	_ = h.db.RecordAuditLog(r.Context(), &domain.AuditLog{
		UserID:    user.ID,
		UserName:  user.Username,
		Action:    "LOGIN",
		Resource:  "auth",
		Details:   "Login berhasil ke dashboard",
		IPAddress: r.RemoteAddr,
	})

	utils.RespondOK(w, map[string]any{
		"success":    true,
		"token":      token,
		"expires_at": exp,
		"user":       user,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	rawToken, _ := r.Context().Value(middleware.TokenContextKey).(string)
	if rawToken != "" {
		claims, err := utils.VerifyJWT(rawToken, h.cfg.JWTSecret)
		exp := time.Now().Add(8 * time.Hour)
		if err == nil && claims.ExpiresAt != nil {
			exp = claims.ExpiresAt.Time
		}
		_ = h.db.RevokeToken(r.Context(), utils.HashTokenSHA256(rawToken), exp)
	}
	if u, ok := middleware.GetAuthUser(r.Context()); ok {
		_ = h.db.RecordAuditLog(r.Context(), &domain.AuditLog{
			UserID:    u.ID,
			UserName:  u.Username,
			Action:    "LOGOUT",
			Resource:  "auth",
			Details:   "Sesi dicabut oleh pengguna",
			IPAddress: r.RemoteAddr,
		})
	}
	utils.RespondOK(w, map[string]any{"success": true, "message": "Logout berhasil"})
}

func (h *AuthHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListUsers(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat daftar pengguna")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": users})
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string      `json:"name"`
		Email    string      `json:"email"`
		Username string      `json:"username"`
		Password string      `json:"password"`
		Role     domain.Role `json:"role"`
		Bidang   string      `json:"bidang"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Format data pengguna tidak valid")
		return
	}
	if strings.TrimSpace(req.Username) == "" || len(req.Password) < 8 {
		utils.RespondError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Username wajib diisi dan password minimal 8 karakter")
		return
	}
	if req.Role != domain.RoleSuperadmin && req.Role != domain.RoleAdmin && req.Role != domain.RoleStaff {
		req.Role = domain.RoleStaff
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "HASH_ERROR", "Gagal memproses password")
		return
	}

	u := &domain.User{
		Name:     strings.TrimSpace(req.Name),
		Email:    strings.TrimSpace(req.Email),
		Username: strings.TrimSpace(req.Username),
		Password: hashed,
		Role:     req.Role,
		Bidang:   strings.TrimSpace(req.Bidang),
		IsActive: true,
	}
	if err := h.db.CreateUser(r.Context(), u); err != nil {
		utils.RespondError(w, http.StatusConflict, "USER_EXISTS", "Username atau email sudah terdaftar")
		return
	}

	utils.RespondCreated(w, map[string]any{"success": true, "data": u})
}

func (h *AuthHandler) ToggleUserStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_ID", "ID pengguna tidak valid")
		return
	}
	var req struct {
		IsActive bool `json:"is_active"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "INVALID_JSON", "Payload status tidak valid")
		return
	}

	if err := h.db.SetUserActive(r.Context(), uint(id64), req.IsActive); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "NOT_FOUND", "Pengguna tidak ditemukan")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memperbarui status pengguna")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true})
}

func (h *AuthHandler) ListAuditLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := h.db.ListAuditLogs(r.Context(), 100)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "DB_ERROR", "Gagal memuat audit log")
		return
	}
	utils.RespondOK(w, map[string]any{"success": true, "data": logs})
}
