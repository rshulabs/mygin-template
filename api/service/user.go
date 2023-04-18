package service

import (
	"backup/global"
	"backup/model"
	"backup/model/system"
	"backup/utils"
	"errors"
	"time"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userSrv *userService) Register(params system.Register) (err error, user model.User) {
	var result = global.App.ConfigDB.Where("mobile=?", params.Mobile).Select("id").First(&model.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	// 注意这里
	timestamps := model.Timestamps{CreateAt: time.Now(), UpdateAt: time.Now()}
	user = model.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password)), Email: params.Email, Timestamps: timestamps}
	err = global.App.ConfigDB.Create(&user).Error
	return
}
