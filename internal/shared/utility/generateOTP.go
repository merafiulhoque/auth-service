package utility

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP() (string, error) {
	const digits = "0123456789"
	otp := make([]byte, 6)

	for i := range otp {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[n.Int64()]
	}

	return string(otp), nil
}
