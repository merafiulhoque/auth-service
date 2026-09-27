package token

import (
	"auth-service/internal/shared/domainerrors"
	"strings"
)

func GetToken(authorization string) (string, error) {
	parts := strings.Split(authorization, " ")

	if len(parts) != 2 {
		return "", domainerrors.TokenErrMalformed
	}

	if parts[0] != "Bearer" {
		return "", domainerrors.TokenErrMalformed
	}

	return parts[1], nil
}
