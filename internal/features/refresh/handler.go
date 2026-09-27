package refresh

import (
	"auth-service/internal/shared/response"
	"net/http"
)

func (h *handler) Refresh() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		refreshToken := r.Header.Get("X-Refresh-Token")

		if refreshToken == "" {
			response.Error(w, http.StatusUnauthorized, "missing refresh token")
			return
		}

		accessToken, newRefreshToken, err := h.service.RefreshService(ctx, refreshToken)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(
			w,
			http.StatusOK,
			"fresh tokens issued",
			map[string]string{
				"access_token":  accessToken,
				"refresh_token": newRefreshToken,
			},
		)
	})
}
