package resultModels

type Result struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 错误码
type ErrorCode int

const (
	SUCCESS     ErrorCode = 200
	FALL        ErrorCode = 400
	NOT_LOGIN   ErrorCode = 411
	VERIFY_FALL ErrorCode = 412
)

const (
	SUCCESS_MSG = "OK"
)

func SuccessResult(data interface{}) Result {
	return Result{SUCCESS, SUCCESS_MSG, data}
}

func ErrorResult(code ErrorCode, message string) Result {
	return Result{code, message, nil}
}
