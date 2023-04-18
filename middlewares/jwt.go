package middlewares

import (
	"backup/global"
	"backup/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			utils.JWTFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(utils.TokenType)+1:]
		// token解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil {
			utils.JWTFail(c)
			c.Abort()
			return
		}
		claims := token.Claims.(*utils.CustomClaims)
		// toekn发布者校验
		if claims.Issuer != GuardName {
			utils.JWTFail(c)
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
