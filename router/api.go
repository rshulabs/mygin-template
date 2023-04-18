package router

import (
	"backup/middlewares"
	"backup/model/system"
	"backup/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/user/register", func(c *gin.Context) {
		var form system.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			utils.ValidatorFail(c, utils.GetErrorMsg(form, err))
			return
		}
		utils.Success(c, "success")
	})
	authRouter := r.Group("").Use(middlewares.JWTAuth(utils.AppGuardName))
	{
		authRouter.POST("/auth/info", func(c *gin.Context) {
			utils.Success(c, "success")
		})
	}
	r.Use(middlewares.Cors())
}
