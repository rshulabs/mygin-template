package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
}
