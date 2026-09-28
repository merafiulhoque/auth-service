package requireuser

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/response"
	"net/http"
)

func (h *handler) RequireUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.Context().Value(domain.UserEmailKey).(string)

		if email == "" {
			response.Error(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		response.Success(w, http.StatusOK, "user fetched successfully", email)
	})
}
