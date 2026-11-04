package middleware

import (
	"net/http"
)

// CORS provides configurable Cross-Origin Resource Sharing middleware
func CORS(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Check origin against whitelist
			isAllowed := false
			for _, o := range allowedOrigins {
				if o == "*" || o == origin {
					isAllowed = true
					break
				}
			}

			if isAllowed && origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else if len(allowedOrigins) > 0 && allowedOrigins[0] == "*" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "86400")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
