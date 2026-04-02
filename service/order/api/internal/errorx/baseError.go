package errorx

const (
	RpcError   = 100000001
	QueryError = 100000002
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

func NewDefaultCodeError() error {
	return &CodeError{
		Code: 10000,
		Msg:  "系统错误",
	}
}

func (e CodeError) Error() string {
	return e.Msg
}

func (e CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
