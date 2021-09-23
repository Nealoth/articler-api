package apierr

var (
	AuthTokenExpired error = NewApiError("User auth token has beed expired", 1)
	AuthTokenInvalid error = NewApiError("User auth token is invalid", 2)
)