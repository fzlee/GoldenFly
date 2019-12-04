package common

const (
	CodeUnauthorized = 1001
	CodePermissionDenied = 1002
)
var ErrorCodes = map[int]string{
	1001: "Unauthorized",
	1002: "Permission Denied",
}
