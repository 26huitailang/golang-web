package response

type ErrorCode struct {
	Code int
	Msg  string
}

var (
	OK              = ErrorCode{1000, "Request success"}
	Error           = ErrorCode{1001, "Request error"}
	ReqParamInvalid = ErrorCode{1002, "Bad request params"}

	AuthFailed        = ErrorCode{2001, "Authentication failed"}
	AuthCookieInvalid = ErrorCode{2002, "Cookie invalid"}
	AuthCookieExpired = ErrorCode{2003, "Cookie expired"}
	AuthLogoutFailed  = ErrorCode{2004, "Logout failed"}
)
