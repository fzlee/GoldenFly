package common

const (
	CodeUnauthorized     = 1001
	CodePermissionDenied = 1002
	CodeNotFound         = 1003
	CodeParameterMissing = 1004

	LoginFailed = 2001
)
var ErrorCodes = map[int]string{
	1001: "Unauthorized",
	1002: "Permission Denied",
	1003: "object not found",
	1004: "certain parameters are missing",
	2001: "Invalid email/password",
}
