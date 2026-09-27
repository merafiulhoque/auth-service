package domainerrors

import "errors"

var (
	ErrEnvLoad = errors.New("failed to load secrets")
)

var (
	ErrDbConnection = errors.New("error connecting database")
)

var (
	ErrNoData       = errors.New("empty body")
	ErrInvalidInput = errors.New("invalid input")
)

var (
	ErrUserAlreadyExists   = errors.New("err user exists")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrUnknownError        = errors.New("unknown error occurred")
	ErrUserNotFound        = errors.New("user not found")
	ErrEntryNotFound       = errors.New("entry not found")
	ErrOtpWrong            = errors.New("otp verification failed")
	ErrOtpExpired          = errors.New("otp expired")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)

var (
	TokenErrMalformed            = errors.New("malformed token")
	TokenErrInvalid              = errors.New("invalid token")
	TokenErrInvalidSigningMethod = errors.New("invalid signing method")
)

var (
	AuthErrUnauthorized = errors.New("unauthorized")
)

var (
	ServerError = errors.New("server error")
)

var (
	SomethingWentWrong = errors.New("something went wrong, please try again after sometime")
)

var (
	RedisErrResetLinkExpired = errors.New("password reset link expired")
)
