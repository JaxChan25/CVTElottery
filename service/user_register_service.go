package service

import (
	"singo/model"
	"singo/serializer"
	"singo/util"
	"time"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=40"`
	RealName        string `form:"real_name" json:"real_name"`
	Avatar          string `form:"avatar" json:"avatar"`
	Mobile          string `form:"mobile" json:"mobile" binding:"required,min=11,max=11"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.GameUser{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {

	user := model.GameUser{
		UserName:      service.UserName,
		RealName:      service.RealName,
		Mobile:        service.Mobile,
		LastLoginTime: time.Now(),
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	//生成16位的salt
	salt := util.RandStringRunes(16)
	// 加密密码
	if err := user.SetPassword(service.Password, salt); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	//设置盐值
	user.Salt = salt

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}
