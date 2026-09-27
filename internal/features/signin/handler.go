package signin

import (
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/validator"
	"encoding/json"
	"net/http"
)

func (h *handler) Signin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data SigninRequest
		ctx := r.Context()

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&data); err != nil {
			response.Error(w, 400, err.Error())
			return
		}

		if err := validator.ValidateStruct(data); err != nil {
			response.Error(w, 422, err.Error())
			return
		}
		accessToken, refreshToken, err := h.s.SigninService(ctx, &data)
		if err != nil {
			response.Error(w, 401, err.Error())
			return
		}

		response.Success(
			w,
			http.StatusOK,
			"login successfull",
			map[string]string{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		)
	})
}
