package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"satpolpp-sulsel-api/internal/core/domain"
)

type contextKey string

const (
	UserContextKey contextKey = "authenticated_user"
)

type AuthUser struct {
	ID       uint        `json:"id"`
	Username string      `json:"username"`
	Role     domain.Role `json:"role"`
}

// RequireAuth validates Bearer token and populates user in request context
func RequireAuth(allowedRoles ...domain.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"success": false,
					"error":   "Akses ditolak: Token autentikasi tidak ditemukan",
					"code":    "UNAUTHORIZED",
				})
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			// In development/mock mode: parse role from token format "jwt-mock-{ROLE}-{ID}"
			var user AuthUser
			if strings.HasPrefix(token, "jwt-mock-") {
				parts := strings.Split(token, "-")
				if len(parts) >= 4 {
					user.Role = domain.Role(parts[2])
					user.Username = "petugas"
					user.ID = 1
				}
			} else {
				// Default staff fallback for local dev
				user.Role = domain.RoleStaff
				user.Username = "petugas"
				user.ID = 1
			}

			// Validate allowed roles if specified
			if len(allowedRoles) > 0 {
				roleAllowed := false
				for _, r := range allowedRoles {
					if user.Role == r || user.Role == domain.RoleSuperadmin {
						roleAllowed = true
						break
					}
				}

				if !roleAllowed {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusForbidden)
					json.NewEncoder(w).Encode(map[string]any{
						"success": false,
						"error":   "Akses ditolak: Anda tidak memiliki wewenang untuk sumber daya ini",
						"code":    "FORBIDDEN",
					})
					return
				}
			}

			ctx := context.WithValue(r.Context(), UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
