package model

import "time"

// ListActivityPrizesResult 该服务器的结果模型
type ListActivityPrizesResult struct {
	UserName   string
	RealName   string `gorm:"type:varchar(50);not null"` //真实姓名
	Mobile     string `gorm:"type:varchar(11);not null"` //真实姓名
	CreatedAt  time.Time
	PrizeName  string
	PrizeLevel string
	Province   string `gorm:"type:varchar(15);not null"`  //省
	City       string `gorm:"type:varchar(15);not null"`  //市
	District   string `gorm:"type:varchar(15);not null"`  //区
	Detail     string `gorm:"type:varchar(100);not null"` //详细地址
}
