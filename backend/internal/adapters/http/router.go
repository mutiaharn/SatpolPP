package httpadapter

import (
	"net/http"
	"time"

	"satpolpp-sulsel-api/internal/adapters/http/handlers"
	"satpolpp-sulsel-api/internal/adapters/middleware"
	"satpolpp-sulsel-api/internal/adapters/repository/postgres"
	"satpolpp-sulsel-api/internal/config"
	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

func NewRouter(
	cfg *config.Config,
	db *postgres.DB,
	authH *handlers.AuthHandler,
	compH *handlers.ComplaintHandler,
	portalH *handlers.PortalHandler,
	uploadH *handlers.UploadHandler,
) http.Handler {
	mux := http.NewServeMux()

	// Healthcheck
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondOK(w, map[string]any{
			"app":       "Satpol PP Sulsel & SIAP TERTIB Unified API",
			"status":    "healthy",
			"version":   "2.0.0",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// Public Portal Routes
	mux.HandleFunc("GET /api/v1/portal/settings", portalH.GetSettings)
	mux.HandleFunc("GET /api/v1/portal/regencies", portalH.ListRegencies)
	mux.HandleFunc("GET /api/v1/portal/articles", portalH.ListArticles)
	mux.HandleFunc("GET /api/v1/portal/articles/{slug}", portalH.GetArticleBySlug)
	mux.HandleFunc("GET /api/v1/portal/documents", portalH.ListDocuments)
	mux.HandleFunc("GET /api/v1/portal/galleries", portalH.ListGalleries)
	mux.HandleFunc("GET /api/v1/portal/organization", portalH.ListOrganizations)

	// Public SIAP TERTIB Routes
	mux.HandleFunc("POST /api/v1/upload/evidence", uploadH.UploadEvidence)
	mux.HandleFunc("POST /api/v1/siaptertib/lapor", compH.SubmitComplaint)
	mux.HandleFunc("POST /api/v1/complaints", compH.SubmitComplaint) // Alias for Nginx rate-limit path
	mux.HandleFunc("GET /api/v1/siaptertib/lacak/{code}", compH.TrackComplaint)

	// Dev Local Uploads Serving (when RustFS S3_ENDPOINT is not configured)
	mux.Handle("GET /api/v1/uploads/", http.StripPrefix("/api/v1/uploads/", http.FileServer(http.Dir(cfg.StoragePath))))

	// Auth Routes
	mux.HandleFunc("POST /api/v1/auth/login", authH.Login)
	mux.Handle("POST /api/v1/auth/logout", middleware.RequireAuth(cfg.JWTSecret, db)(http.HandlerFunc(authH.Logout)))

	// Helper wrappers for RBAC
	staffOrAbove := func(fn http.HandlerFunc) http.Handler {
		return middleware.RequireAuth(cfg.JWTSecret, db, domain.RoleStaff, domain.RoleAdmin, domain.RoleSuperadmin)(fn)
	}
	adminOrAbove := func(fn http.HandlerFunc) http.Handler {
		return middleware.RequireAuth(cfg.JWTSecret, db, domain.RoleAdmin, domain.RoleSuperadmin)(fn)
	}
	superadminOnly := func(fn http.HandlerFunc) http.Handler {
		return middleware.RequireAuth(cfg.JWTSecret, db, domain.RoleSuperadmin)(fn)
	}

	// Admin & Staff Routes (Complaints + CMS)
	mux.Handle("GET /api/v1/admin/complaints", staffOrAbove(compH.ListAdminComplaints))
	mux.Handle("PUT /api/v1/admin/complaints/{ticket}/status", staffOrAbove(compH.UpdateComplaintStatus))
	mux.Handle("POST /api/v1/admin/complaints/{ticket}/unmask", adminOrAbove(compH.UnmaskComplaintIdentity))
	mux.Handle("GET /api/v1/admin/evidence/", staffOrAbove(compH.StreamPrivateEvidence))

	mux.Handle("POST /api/v1/admin/upload/media", staffOrAbove(uploadH.UploadMedia))
	mux.Handle("POST /api/v1/admin/upload/document", staffOrAbove(uploadH.UploadDocument))

	mux.Handle("POST /api/v1/admin/articles", staffOrAbove(portalH.CreateArticle))
	mux.Handle("DELETE /api/v1/admin/articles/{id}", adminOrAbove(portalH.DeleteArticle))

	mux.Handle("POST /api/v1/admin/documents", staffOrAbove(portalH.CreateDocument))
	mux.Handle("DELETE /api/v1/admin/documents/{id}", adminOrAbove(portalH.DeleteDocument))

	mux.Handle("POST /api/v1/admin/galleries", staffOrAbove(portalH.CreateGallery))
	mux.Handle("DELETE /api/v1/admin/galleries/{id}", adminOrAbove(portalH.DeleteGallery))

	mux.Handle("POST /api/v1/admin/organization", adminOrAbove(portalH.UpsertOrganization))
	mux.Handle("PUT /api/v1/admin/organization/{id}", adminOrAbove(portalH.UpsertOrganization))
	mux.Handle("DELETE /api/v1/admin/organization/{id}", adminOrAbove(portalH.DeleteOrganization))

	// Superadmin Routes
	mux.Handle("PUT /api/v1/superadmin/settings", superadminOnly(portalH.UpdateSettings))
	mux.Handle("GET /api/v1/superadmin/users", superadminOnly(authH.ListUsers))
	mux.Handle("POST /api/v1/superadmin/users", superadminOnly(authH.CreateUser))
	mux.Handle("PATCH /api/v1/superadmin/users/{id}/status", superadminOnly(authH.ToggleUserStatus))
	mux.Handle("GET /api/v1/superadmin/audit-logs", adminOrAbove(authH.ListAuditLogs))

	// Frontier Security Middleware Chain
	rateLimiter := middleware.NewIPRateLimiter(120, time.Second)
	var handler http.Handler = mux
	handler = middleware.SecurityHeaders(handler)
	handler = middleware.CORS([]string{cfg.PublicSiteURL, "http://localhost:3000", "http://localhost:5173"})(handler)
	handler = rateLimiter.Middleware(handler)
	handler = middleware.PanicRecovery(handler)

	return handler
}
