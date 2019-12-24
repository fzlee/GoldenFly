package common

const (
	CodeUnauthorized     = 1001
	CodePermissionDenied = 1002
	CodeNotFound         = 1003
	CodeParameterMissing = 1004

	CodeLoginFailed      = 2001
	CodeFileUploadFailed = 3001
	CodeInvalidPassword  = 3002
	CodeCommentNotFound  = 4001
)

var ErrorCodes = map[int]string{
	1001: "Unauthorized",
	1002: "Permission Denied",
	1003: "object not found",
	1004: "certain parameters are invalid or missing",
	2001: "Invalid email/password",
	3001: "File uploading failed",
	3002: "Invalid Password for Article",
	4001: "parent comment not found",
}
