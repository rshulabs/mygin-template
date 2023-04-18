package router

import (
	"backup/model/system"
	"backup/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/register", func(c *gin.Context) {
		var form system.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": utils.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}
