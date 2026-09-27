package signup

import (
	"auth-service/internal/shared/domainerrors"
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/validator"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *handler) Signup() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var data SignupRequest
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&data); err != nil {
			response.Error(w, 409, err.Error())
			return
		}

		if err := validator.ValidateStruct(data); err != nil {
			response.Error(w, 422, err.Error())
			return
		}
		id, err := h.s.SignupService(ctx, &data)

		if err != nil {

			if errors.Is(err, domainerrors.ErrUserAlreadyExists) {
				response.Error(w, 409, err.Error())
				return
			}

			response.Error(w, 500, err.Error())
			return
		}
		response.Success(w, 201, "user created", id)
	})
}
