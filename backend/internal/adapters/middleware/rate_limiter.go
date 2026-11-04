package middleware

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type clientRecord struct {
	tokens     int
	lastRefill time.Time
}

// IPRateLimiter provides thread-safe IP-based token-bucket rate limiting
type IPRateLimiter struct {
	mu       sync.Mutex
	clients  map[string]*clientRecord
	limit    int           // Max burst tokens
	refill   time.Duration // Time to refill one token
}

// NewIPRateLimiter creates a new rate limiter (e.g. 60 requests per minute)
func NewIPRateLimiter(limit int, refillRate time.Duration) *IPRateLimiter {
	limiter := &IPRateLimiter{
		clients: make(map[string]*clientRecord),
		limit:   limit,
		refill:  refillRate,
	}

	// Background worker to purge stale IPs every 5 minutes
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			limiter.mu.Lock()
			now := time.Now()
			for ip, record := range limiter.clients {
				if now.Sub(record.lastRefill) > 10*time.Minute {
					delete(limiter.clients, ip)
				}
			}
			limiter.mu.Unlock()
		}
	}()

	return limiter
}

func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}

	// Check X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return strings.TrimSpace(realIP)
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// Allow checks if the client IP has available tokens
func (rl *IPRateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	record, exists := rl.clients[ip]
	if !exists {
		rl.clients[ip] = &clientRecord{
			tokens:     rl.limit - 1,
			lastRefill: now,
		}
		return true
	}

	// Calculate tokens refilled since last visit
	elapsed := now.Sub(record.lastRefill)
	refillCount := int(elapsed / rl.refill)
	if refillCount > 0 {
		record.tokens += refillCount
		if record.tokens > rl.limit {
			record.tokens = rl.limit
		}
		record.lastRefill = now
	}

	if record.tokens > 0 {
		record.tokens--
		return true
	}

	return false
}

// Middleware wraps an http.Handler with rate limiting
func (rl *IPRateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getClientIP(r)

		if !rl.Allow(ip) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Retry-After", "60")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"error":   "Terlalu banyak permintaan (Rate limit exceeded). Silakan coba lagi dalam beberapa saat.",
				"code":    "RATE_LIMIT_EXCEEDED",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
