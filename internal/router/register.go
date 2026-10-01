package router

import (
	"auth-service/internal/config"
	"auth-service/internal/features/refresh"
	requireuser "auth-service/internal/features/requireUser"
	"auth-service/internal/features/resetpassword"
	sendotp "auth-service/internal/features/send_otp"
	"auth-service/internal/features/signin"
	"auth-service/internal/features/signout"
	"auth-service/internal/features/signup"
	updatepassword "auth-service/internal/features/updatePassword"
	verifyotp "auth-service/internal/features/verify-otp"
	"auth-service/internal/features/welcome"
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/middleware"
	"auth-service/internal/shared/otp"
	"database/sql"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v3"
)

func RegisterRouter(
	mux *http.ServeMux,
	db *sql.DB,
	cfg *config.Config,
	rdb *redis.Client,
	emailSender *resend.Client,
	otpStore *otp.Store,
) {
	//welcome home route
	rateLimitedWelcomeHandler := middleware.
		RateLimiter(rdb, 20, 1*time.Minute)(
		welcome.CreateNewHandler().
			GETWelcomeAPI())
	mux.Handle("GET /", rateLimitedWelcomeHandler)

	// signup
	signupHandlerWithLimit := middleware.
		RateLimiter(rdb, 2, 1*time.Minute)(
		signup.
			CreateNewHandler(db).
			Signup())
	mux.Handle(domain.POST_SIGNUP, signupHandlerWithLimit)

	//signin
	rateLimitedSigninHandler := middleware.
		AuthMiddleware(cfg.JwtSecret)(
		signin.CreateNewHandler(db, cfg.JwtSecret, rdb).
			Signin())

	mux.Handle(domain.POST_SIGNIN, rateLimitedSigninHandler)

	//send otp
	rateLimitedSendOtp := middleware.
		RateLimiter(rdb, 3, 1*time.Minute)(
		sendotp.CreateNewHandler(db, emailSender, rdb).
			SendOTP(),
	)
	mux.Handle(domain.POST_SEND_OTP, rateLimitedSendOtp)

	// get user
	rateLimitedGetUser := middleware.
		RateLimiter(rdb, 30, 1*time.Minute)(
		middleware.AuthMiddleware(cfg.JwtSecret)(
			requireuser.CreateNewHandler().
				RequireUser(),
		),
	)
	mux.Handle(domain.GET_ME, rateLimitedGetUser)

	//signout
	signout := signout.CreateNewHandler(rdb)
	mux.Handle(
		domain.POST_SIGNOUT,
		middleware.AuthMiddleware(cfg.JwtSecret)(signout.Signout()),
	)

	//verify otp
	rateLimitedVerifyOtp := middleware.RateLimiter(rdb, 6, 1*time.Minute)(
		verifyotp.CreateNewHandler(db, rdb, otpStore).
			VerifyOTP(),
	)
	mux.Handle(domain.POST_VERIFY_OTP, rateLimitedVerifyOtp)

	// refresh
	rateLimitedRefresh := middleware.RateLimiter(rdb, 5, 1*time.Minute)(
		refresh.CreateNewHandler(rdb, cfg.JwtSecret).
			Refresh(),
	)
	mux.Handle(domain.GET_REFRESH, rateLimitedRefresh)

	//reset-password -- get reset link
	rateLimitedResetPassword := middleware.RateLimiter(rdb, 1, 1*time.Minute)(
		resetpassword.CreateNewHandler(db, rdb, emailSender, cfg.AllowedOrigin).
			ResetPassword(),
	)
	mux.Handle(domain.POST_RESET_PASSWORD, rateLimitedResetPassword)

	// update password via link
	rateLimitedUpdatePassword := middleware.RateLimiter(rdb, 2, 1*time.Minute)(
		updatepassword.CreateNewHandler(db, rdb).
			UpdatePassword(),
	)
	mux.Handle(domain.POST_UPDATE_PASSWORD, rateLimitedUpdatePassword)
}
