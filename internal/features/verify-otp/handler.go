package verifyotp

import (
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/validator"
	"encoding/json"
	"net/http"
)

func (h *handler) VerifyOTP() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var data VerifyOTP

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&data); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := validator.ValidateStruct(data); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if err := h.service.VerifyOTPService(ctx, data.Email, data.Otp); err != nil {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Success(w, http.StatusOK, "verification success", nil)
	})
}
