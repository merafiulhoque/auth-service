package domain

const (
	ConstRedisOtpKey          = "EMAIL_VERIFICATION_OTP:"
	ConstRedisRefreshTokenKey = "REFRESH_TOKEN:"
	ConstRedisResetLinkKey    = "PASSWORD_RESET_URL:"
	ConstRateLimitKey         = "RATE_LIMIT:"
)

type UserEmail string

const UserEmailKey UserEmail = "email"
