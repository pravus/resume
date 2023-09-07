package model

type Error interface {
	error
	Code() ErrorCode
	Cause() error
}

type ErrorCode string

type _error struct {
	code  ErrorCode
	cause error
}

var _ Error = (*_error)(nil)

func NewError(code ErrorCode, cause error) *_error {
	var err = &_error{
		code:  code,
		cause: cause,
	}
	return err
}

func (err _error) Code() ErrorCode {
	return err.code
}

func (err _error) Cause() error {
	return err.cause
}

func (err _error) Error() string {
	if err.cause == nil {
		return string(err.code)
	} else {
		return string(err.code) + `: ` + err.cause.Error()
	}
}
