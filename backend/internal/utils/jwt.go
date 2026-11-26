package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"satpolpp-sulsel-api/internal/core/domain"
)

type JWTClaims struct {
	UserID   uint        `json:"uid"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
	Role     domain.Role `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT issues a signed HS256 JWT token valid for the specified duration.
func GenerateJWT(user *domain.User, secret string, ttl time.Duration) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(ttl)

	claims := JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Name,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "satpolpp-sulsel-api",
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	return signed, exp, err
}

// VerifyJWT parses and validates an HS256 JWT signature and expiration.
func VerifyJWT(tokenStr, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
