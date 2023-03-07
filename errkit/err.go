package errkit

type Err struct {
	Code    Code
	Message string
}

func (e *Err) Error() string {
	return e.Message
}

func (e *Err) CodeToHttpStatus() int {
	code := e.Code
	if code >= _MaxSize {
		code = _MaxSize
	}

	return Code2HttpStatus[e.Code]
}

func NewErr(code Code, msg string) *Err {
	return &Err{
		Code:    code,
		Message: msg,
	}
}
