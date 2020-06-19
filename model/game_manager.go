package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GameManager 活动管理员
type GameManager struct {
	gorm.Model
	Activities    []Activity //应该手握很多活动
	UserName      string     `gorm:"type:varchar(50);not null"`
	Password      string     `gorm:"type:varchar(50);not null"`
	Avatar        string     `gorm:"default:'default.png'"`
	Moblie        string     `gorm:"type:varchar(11);not null"`
	LastLoginTime time.Time
}
