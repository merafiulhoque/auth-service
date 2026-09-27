package validator

import (
	"auth-service/internal/shared/domainerrors"
	"errors"
	"io"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s any) error {
	err := validate.Struct(s)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return domainerrors.ErrNoData
		}
		return domainerrors.ErrInvalidInput
	}
	return nil
}
