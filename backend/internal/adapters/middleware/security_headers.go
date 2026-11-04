package middleware

import (
	"net/http"
)

// SecurityHeaders applies defense-in-depth HTTP security headers suitable for government portals
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent MIME-sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Prevent Clickjacking (iframe embedding)
		w.Header().Set("X-Frame-Options", "DENY")

		// Legacy XSS Protection
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// Referrer Policy: Send origin when cross-origin
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		// Restrict dangerous browser features
		w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=(self)")

		// HTTP Strict Transport Security (HSTS)
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		// Strict Content Security Policy (CSP)
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https: blob:; font-src 'self' data:; connect-src 'self' http://localhost:* https://*.sulselprov.go.id; frame-ancestors 'none';")

		next.ServeHTTP(w, r)
	})
}
