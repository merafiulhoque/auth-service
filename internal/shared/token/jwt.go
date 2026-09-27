package token

import (
	"auth-service/internal/shared/domainerrors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(email string, secret string) (string, string, error) {
	accessClaims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth-service",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Subject:   "auth",
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(secret))

	if err != nil {
		return "", "", err
	}

	refreshClaims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth-service",
			Subject:   "auth",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret))

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func VerifyToken(token string, secret string) (string, error) {
	claims := &Claims{}

	parsedToken, err := jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", domainerrors.TokenErrInvalidSigningMethod
			}

			return []byte(secret), nil
		},
	)

	if parsedToken == nil || !parsedToken.Valid {
		return "", domainerrors.TokenErrInvalid
	}

	if err != nil {
		return "", err
	}

	return claims.Email, nil
}
