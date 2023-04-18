package utils

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
