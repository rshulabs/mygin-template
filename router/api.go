package router

import (
	"backup/api/controller"
	"backup/middlewares"
	"backup/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/user/register", controller.Register)
	authRouter := r.Group("").Use(middlewares.JWTAuth(utils.AppGuardName))
	{
		authRouter.POST("/auth/info", func(c *gin.Context) {
			utils.Success(c, "success")
		})
	}
	r.Use(middlewares.Cors())
}
