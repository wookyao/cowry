package errkit

import "net/http"

type Code uint32

const (
	// OK success
	OK Code = iota

	// Canceled 取消
	Canceled

	// UnKnown 未知
	UnKnown

	// InvalidArgument 参数错误
	InvalidArgument

	// NotFound 不存在
	NotFound

	// AlreadyExists 已存在
	AlreadyExists

	// PermissionDenied 没有权限；拒绝访问
	PermissionDenied

	// Unauthenticated 未验证，验证失败
	Unauthenticated

	// TimeOut 操作超时
	TimeOut

	// BadRequest 错误请求
	BadRequest

	// Internal 内部错误
	Internal

	// Unavailable 不可用的
	Unavailable

	// DataLoss 数据丢失
	DataLoss

	// DBError 服务器错误
	DBError

	// MissMatch 错误匹配
	MissMatch

	// _MaxSize 最大值范围
	_MaxSize
)

var Code2HttpStatus = map[Code]int{
	OK:               http.StatusOK,
	Canceled:         http.StatusNotImplemented,
	UnKnown:          http.StatusInternalServerError,
	InvalidArgument:  http.StatusBadRequest,
	NotFound:         http.StatusNotFound,
	AlreadyExists:    http.StatusInternalServerError,
	PermissionDenied: http.StatusForbidden,
	Unauthenticated:  http.StatusUnauthorized,
	TimeOut:          http.StatusRequestTimeout,
	BadRequest:       http.StatusBadRequest,
	Internal:         http.StatusInternalServerError,
	Unavailable:      http.StatusServiceUnavailable,
	DataLoss:         http.StatusInternalServerError,
	DBError:          http.StatusInternalServerError,
	MissMatch:        http.StatusInternalServerError,
	_MaxSize:         http.StatusInternalServerError,
}
