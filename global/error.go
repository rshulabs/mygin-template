package global

/*
错误码
*/
type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}
type CustomErrors struct {
	BusinessError  CustomError
	ValidatorError CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{
		40000, "业务错误",
	},
	ValidatorError: CustomError{
		42000, "请求参数错误",
	},
}
