package updatepassword

import (
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/validator"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type UpdatePasswordReq struct {
	Password string `json:"password" validate:"required"`
}

func (h *handler) UpdatePassword() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data UpdatePasswordReq

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&data); err != nil {
			if errors.Is(err, io.EOF) {
				response.Error(w, http.StatusBadRequest, "empty request body")
				return
			}
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := validator.ValidateStruct(data); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ctx := r.Context()
		token := r.URL.Query().Get("token")
		err := h.service.UpdatepasswordService(ctx, token, data.Password)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(w, http.StatusOK, "password reset successfull", nil)
	})
}
