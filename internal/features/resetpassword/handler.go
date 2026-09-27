package resetpassword

import (
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/validator"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (h *handler) ResetPassword() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data ResetPasswordRequest

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&data); err != nil {
			if errors.Is(err, io.EOF) {
				response.Error(w, http.StatusBadRequest, "no data in request body")
				return
			}
			response.Error(w, http.StatusBadRequest, "invalid request")
			return
		}

		if err := validator.ValidateStruct(data); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "invalid input")
			return
		}
		ctx := r.Context()

		err := h.service.ResetPasswordService(ctx, data.Email)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, "something went wrong")
			return
		}

		response.Success(w, http.StatusOK, "password reset link sent to your email", nil)

	})
}

