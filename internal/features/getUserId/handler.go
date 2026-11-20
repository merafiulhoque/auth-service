package getuserid

import (
	"auth-service/internal/shared/domainerrors"
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/validator"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (h *handler) GetUserId() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var data GetUserId

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&data); err != nil {
			if errors.Is(err, io.EOF) {
				response.Error(w, http.StatusBadRequest, domainerrors.ErrNoData.Error())
				return
			}
			response.Error(w, http.StatusBadRequest, domainerrors.ErrInvalidInput.Error())
			return
		}

		if err := validator.ValidateStruct(data); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, domainerrors.ErrInvalidInput.Error())
			return
		}

		id, err := h.service.getUserIdService(ctx, data.Email)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(w, http.StatusOK, "id fetched successfully", id)
	})
}
