package system

import "backup/utils"

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (r Register) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号格式不对",
		"password.required": "用户密码不能为空",
	}
}
