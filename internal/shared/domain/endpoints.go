package domain

const (
	version              = "v1"
	POST_SIGNUP          = "POST /api/" + version + "/auth/signup"
	POST_SIGNIN          = "POST /api/" + version + "/auth/signin"
	POST_SIGNOUT         = "POST /api/" + version + "/auth/signout"
	POST_SEND_OTP        = "POST /api/" + version + "/auth/send-otp"
	POST_VERIFY_OTP      = "POST /api/" + version + "/auth/verify-otp"
	GET_REFRESH          = "GET /api/" + version + "/auth/refresh"
	POST_RESET_PASSWORD  = "POST /api/" + version + "/auth/reset-password"
	POST_UPDATE_PASSWORD = "POST /api/" + version + "/auth/update-password"
	GET_ME               = "GET /api/" + version + "/auth/me"
)
