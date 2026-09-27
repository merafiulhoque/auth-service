package sendotp

import (
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/validator"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type sendotp struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *handler) SendOTP() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		defer ctx.Done()
		var data sendotp
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&data); err != nil {
			if errors.Is(err, io.EOF) {
				response.Error(w, http.StatusBadRequest, "no data in requet body")
				return
			}
			response.Error(w, http.StatusBadRequest, "bad request")
			return
		}

		if err := validator.ValidateStruct(data); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		err := h.s.SendOTPService(ctx, data)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		response.Success(w, http.StatusOK, "email sent", nil)

	})
}
