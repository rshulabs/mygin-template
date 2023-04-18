package initialize

import (
	"backup/global"
	"backup/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

func OtherInit() {
	initializeValidator()
	initRedis()
	fmt.Println(" ===== Other init ===== ")
}

// 验证器注册
func initializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)
		_ = v.RegisterValidation("email", utils.ValidatorEmail)
		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// Redis redis初始化
func initRedis() {
	redisCfg := global.App.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.ConfigZap.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println(" redis init success")
		global.App.ConfigZap.Info("redis connect ping response:", zap.String("pong", pong))
		global.App.Redis = client
	}
}
