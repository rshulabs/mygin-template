package controller

import (
	"backup/api/service"
	"backup/model/system"
	"backup/utils"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form system.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.ValidatorFail(c, utils.GetErrorMsg(form, err))
		return
	}
	if err, user := service.UserService.Register(form); err != nil {
		utils.BusinessFail(c, err.Error())
	} else {
		utils.Success(c, user)
	}
}
