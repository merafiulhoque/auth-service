package signout

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"auth-service/internal/shared/response"
	"net/http"
)

func (h *handler) Signout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		email, ok := ctx.Value(domain.UserEmailKey).(string)

		if !ok || email == "" {
			response.Error(w, http.StatusUnauthorized, domainerrors.AuthErrUnauthorized.Error())
			return
		}

		err := h.service.SignoutService(ctx, email)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(w, http.StatusOK, "signout successfull", nil)
	})
}
