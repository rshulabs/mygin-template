package utils

import (
	"backup/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtService struct {
}

var JwtService = new(jwtService)

// JwtUser 所有需要颁发token的用户模型必须实现此接口
type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义claims
type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type TokenOutput struct {
	AccessToken string `yaml:"access_token"`
	ExpiresIn   int    `json:"expires-in"`
	TokenType   string `json:"token-type"`
}

// CreateToken 生成token
func (js *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutput, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.Expire,
				Id:        user.GetUid(),
				Issuer:    GuardName, // 用于在中间件区分不同客户端发的token，避免token跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))
	tokenData = TokenOutput{
		tokenStr,
		int(global.App.Config.Jwt.Expire),
		TokenType,
	}
	return
}
