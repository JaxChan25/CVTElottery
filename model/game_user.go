package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// GameUser 参与游戏用户-模型
type GameUser struct {
	gorm.Model
	Address        Address      //包含一个地址
	UserActions    []UserAction //包含多个Action
	PasswordDigest string       `gorm:"not null"`                  //密码
	UserName       string       `gorm:"type:varchar(50);not null"` //用户登录名
	RealName       string       `gorm:"type:varchar(50);not null"` //真实姓名
	Avatar         string       `gorm:"default:'default.png'"`     //头像，可以为null
	Mobile         string       `gorm:"type:varchar(11);not null"` //电话号码
	LastLoginTime  time.Time
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (GameUser, error) {
	var user GameUser
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *GameUser) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *GameUser) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
