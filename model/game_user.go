package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/jinzhu/gorm"
)

// GameUser 参与游戏用户-模型
type GameUser struct {
	gorm.Model
	Address        Address      //包含一个地址
	UserActions    []UserAction //包含多个Action
	PasswordDigest string       `gorm:"not null"`                  //密码
	UserName       string       `gorm:"type:varchar(50);not null"` //用户登录名
	RealName       string       `gorm:"type:varchar(50)"`          //真实姓名
	Avatar         string       `gorm:"default:'default.png'"`     //头像，可以为null
	Mobile         string       `gorm:"type:varchar(11);not null"` //电话号码
	Salt           string       `gorm:"type:varchar(16);not null"`
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

//Generatemd5 生成md5
func Generatemd5(str string, salt string) string {
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

// SetPassword 设置密码
func (user *GameUser) SetPassword(password string, salt string) error {
	bytes := Generatemd5(password, salt)
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *GameUser) CheckPassword(password string) bool {
	return user.PasswordDigest == Generatemd5(password, user.Salt)
}
