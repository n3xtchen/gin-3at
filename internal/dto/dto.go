package dto

type Code int

const (
	CodeSuccess       Code = 0
	CodeParamError         = 10001
	CodeUnauth             = 20001
	CodePermission         = 20002
	CodeResourceExist      = 30001
	CodeResourceMiss       = 30002
	CodeDBError            = 40001
	CodeUnknownError       = 50000
)

type APIResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}
