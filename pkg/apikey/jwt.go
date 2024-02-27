package apikey

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type Claims struct {
	APIKey    string `json:"api_key"`
	Timestamp int64  `json:"timestamp"`
	jwt.RegisteredClaims
}

func CreateToken(apiKey string) string {
	parts := strings.Split(apiKey, ".")
	if len(parts) != 2 {
		return ""
	}
	id, secret := parts[0], parts[1]
	claims := Claims{
		APIKey:    id,
		Timestamp: time.Now().UnixMilli(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(600) * time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["alg"] = "HS256"
	token.Header["sign_type"] = "SIGN"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return err.Error()
	}
	return tokenString
}
