package middleware

import (
	"context"
	"net/http"
	"strings"

	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

type contextKey string

const (
	UserContextKey  contextKey = "authenticated_user"
	TokenContextKey contextKey = "raw_jwt_token"
)

type AuthUser struct {
	ID       uint        `json:"id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
	Role     domain.Role `json:"role"`
}

// TokenRevocationChecker checks if a token hash has been revoked (logout/suspend).
type TokenRevocationChecker interface {
	IsTokenRevoked(ctx context.Context, tokenHash string) (bool, error)
}

// RequireAuth validates real HS256 JWT tokens, checks revocation, and enforces RBAC roles.
func RequireAuth(jwtSecret string, revoker TokenRevocationChecker, allowedRoles ...domain.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				utils.RespondError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Akses ditolak: Token autentikasi tidak ditemukan")
				return
			}

			tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			claims, err := utils.VerifyJWT(tokenStr, jwtSecret)
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, "INVALID_TOKEN", "Sesi tidak valid atau telah kedaluwarsa")
				return
			}

			if revoker != nil {
				revoked, err := revoker.IsTokenRevoked(r.Context(), utils.HashTokenSHA256(tokenStr))
				if err != nil || revoked {
					utils.RespondError(w, http.StatusUnauthorized, "TOKEN_REVOKED", "Sesi telah dicabut, silakan login kembali")
					return
				}
			}

			if len(allowedRoles) > 0 {
				roleAllowed := false
				for _, allowed := range allowedRoles {
					if claims.Role == allowed || claims.Role == domain.RoleSuperadmin {
						roleAllowed = true
						break
					}
				}
				if !roleAllowed {
					utils.RespondError(w, http.StatusForbidden, "FORBIDDEN", "Akses ditolak: Anda tidak memiliki wewenang untuk sumber daya ini")
					return
				}
			}

			user := AuthUser{
				ID:       claims.UserID,
				Username: claims.Username,
				Name:     claims.Name,
				Role:     claims.Role,
			}

			ctx := context.WithValue(r.Context(), UserContextKey, user)
			ctx = context.WithValue(ctx, TokenContextKey, tokenStr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetAuthUser extracts the authenticated user from request context.
func GetAuthUser(ctx context.Context) (AuthUser, bool) {
	u, ok := ctx.Value(UserContextKey).(AuthUser)
	return u, ok
}
