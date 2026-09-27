package router

import (
	"auth-service/internal/config"
	"auth-service/internal/features/refresh"
	"auth-service/internal/features/resetpassword"
	sendotp "auth-service/internal/features/send_otp"
	"auth-service/internal/features/signin"
	"auth-service/internal/features/signout"
	"auth-service/internal/features/signup"
	updatepassword "auth-service/internal/features/updatePassword"
	verifyotp "auth-service/internal/features/verify-otp"
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/middleware"
	"auth-service/internal/shared/otp"
	"database/sql"
	"net/http"

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
	// signup
	signupHandler := signup.CreateNewHandler(db)
	mux.Handle(domain.POST_SIGNUP, signupHandler.Signup())

	//signin
	signinHandler := signin.CreateNewHandler(db, cfg.JwtSecret, rdb)
	mux.Handle(domain.POST_SIGNIN, signinHandler.Signin())

	//send otp
	sendOtp := sendotp.CreateNewHandler(db, emailSender, rdb)
	mux.Handle(domain.POST_SEND_OTP, sendOtp.SendOTP())

	//signout
	signout := signout.CreateNewHandler(rdb)
	mux.Handle(
		domain.POST_SIGNOUT,
		middleware.AuthMiddleware(
			signout.Signout(),
			cfg.JwtSecret,
		),
	)

	//verify otp
	verifyOtp := verifyotp.CreateNewHandler(db, rdb, otpStore)
	mux.Handle(domain.POST_VERIFY_OTP, verifyOtp.VerifyOTP())

	// refresh
	refresh := refresh.CreateNewHandler(rdb, cfg.JwtSecret)
	mux.Handle(domain.GET_REFRESH, refresh.Refresh())

	//reset-password -- get reset link
	resetPasswordLink := resetpassword.CreateNewHandler(db, rdb, emailSender, cfg.AllowedOrigin)
	mux.Handle(domain.POST_RESET_PASSWORD, resetPasswordLink.ResetPassword())

	// update password via link
	updatePassword := updatepassword.CreateNewHandler(db, rdb)
	mux.Handle(domain.POST_UPDATE_PASSWORD, updatePassword.UpdatePassword())
}
