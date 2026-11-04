package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"runtime/debug"
)

// PanicRecovery recovers from any runtime panics and returns a clean 500 error
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				stack := string(debug.Stack())
				slog.Error("Internal Server Panic Caught",
					"error", rec,
					"path", r.URL.Path,
					"method", r.Method,
					"stack", stack,
				)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"success": false,
					"error":   "Terjadi kesalahan internal pada server. Tim pengembang telah diberitahu.",
					"code":    "INTERNAL_SERVER_ERROR",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
