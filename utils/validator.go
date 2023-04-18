package utils

/*
验证器模块
1.定义一个验证器规则接口，实现一个规则map方法
2.定义验证错误信息方法，得到错误验证
3.扩展自定义验证器，启动初始化
*/
import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type ValidatorMessages map[string]string
type Validator interface {
	GetMessages() ValidatorMessages
}

func GetErrorMsg(req interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := req.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			if isValidator {
				if msg, exist := req.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return msg
				}
			}
		}
		return err.Error()
	}
	return "parameter error"
}

// ValidateMobile 校验mobile
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}

// ValidatorEmail 验证email
func ValidatorEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9.!#$%&’*+/=?^_`+`){|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`, email)
	if !ok {
		return false
	}
	return true
}
