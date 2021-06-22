package errors

import (
	"errors"
)

var (
	codes = make(map[error]ErrorCode)
)

// ErrorCode 错误码
type ErrorCode struct {
	Code           int
	Message        string
	HTTPStatusCode int
}

// newErrorCode 设定错误码
func newErrorCode(err error, code int, message string, status ...int) error {
	errCode := ErrorCode{
		Code:    code,
		Message: message,
	}
	if len(status) > 0 {
		errCode.HTTPStatusCode = status[0]
	}
	codes[err] = errCode
	return err
}

// FromErrorCode 获取错误码
func FromErrorCode(err error) (ErrorCode, bool) {
	v, ok := codes[err]
	return v, ok
}

// newBadRequestError 创建请求错误
func newBadRequestError(err error) {
	newErrorCode(err, 400, err.Error(), 400)
}

// New400Error new bad request error
//func New400Error(msg string) error {
//	err := errors.New(msg)
//	newBadRequestError(err)
//	return err
//}

// New400Error new bad request error
func New400Error(msg string) error {
	err := errors.New(msg)
	var found bool
	for k := range codes {
		if k.Error() == msg {
			found = true
			err = k
		}
	}
	if !found {
		newBadRequestError(err)
	}
	return err
}

// newUnauthorizedError 创建未授权错误
func newUnauthorizedError(err error) {
	newErrorCode(err, 401, err.Error(), 401)
}

// newInternalServerError 创建服务器错误
func newInternalServerError(err error) {
	newErrorCode(err, 500, err.Error(), 500)
}

// newDBServerInternalError 创建数据库错误
func newDBServerInternalError(err error) {
	newErrorCode(err, 50001, "数据库发生错误", 500)
}
