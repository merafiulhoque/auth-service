package token

import (
	"auth-service/internal/shared/domainerrors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, email string, secret string) (string, string, error) {
	accessClaims := Claims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth-service",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 60 * time.Minute)),
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

func VerifyToken(token string, secret string) (int, string, error) {
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
		return 0, "", domainerrors.TokenErrInvalid
	}

	if err != nil {
		return 0, "", err
	}

	return claims.UserID, claims.Email, nil
}
