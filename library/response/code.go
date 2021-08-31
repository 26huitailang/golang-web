package response

const (
	SystemError     int = -1
	OK              int = 0
	ReqParamInvalid int = 1
)
const (
	AUTH = iota + 10000
	AuthFailed
	AuthCreateSessionFailed
	AuthCookieInvalid
	AuthCookieExpired
	AuthLogoutFailed
)
