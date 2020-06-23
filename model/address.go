package model

import (
	"github.com/jinzhu/gorm"
)

// Address 地址模型
type Address struct {
	gorm.Model
	GameUserID uint   //外码
	RealName   string `gorm:"type:varchar(50);not null"`  //真实姓名
	Mobile     string `gorm:"type:varchar(11);not null"`  //电话号码
	Province   string `gorm:"type:varchar(15);not null"`  //省
	City       string `gorm:"type:varchar(15);not null"`  //市
	District   string `gorm:"type:varchar(15);not null"`  //区
	Detail     string `gorm:"type:varchar(100);not null"` //详细地址
}
