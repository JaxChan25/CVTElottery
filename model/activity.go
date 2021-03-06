package model

import (
	"singo/cache"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// Activity 活动模型
type Activity struct {
	gorm.Model
	UserActions     []UserAction //包含多个Action
	GamePrizes      []GamePrize  //包含多个奖品
	GameManagerID   uint         //外码
	Name            string       `gorm:"type:varchar(15);not null"`
	Title           string       `gorm:"type:varchar(20);default:'大奖盘抽奖';not null"`
	StartTime       time.Time    `gorm:"not null"`
	EndTime         time.Time
	LimitType       int    `gorm:"not null"`                   //(0:无限制,1:每日抽奖次数限制,2:总抽奖次数限制)',
	LimitNum        int    `gorm:"not null;default:'0'"`       // '限制的抽奖次数',
	RuleText        string `gorm:"type:varchar(512);not null"` // '活动规则介绍'
	BannerImage     string `gorm:"type:varchar(150);not null;default:'BannerImager_default.png'"`
	LotteryImage    string `gorm:"type:varchar(150);not null;default:'LotteryImage_default.png'"`
	BackgrouldColor string `gorm:"type:varchar(20);not null;default:'#FFFFFF'"`
	VirtualNum      int    `gorm:"not null"` //'虚拟参与者个数'
}

// //Addview 活动浏览+1
// func (activity *Activity) Addview() {
// 	cache.RedisClient.Incr(cache.ActivityViewKey(activity.ID))
// }

// //ViewNumber 游览累计数
// func (activity *Activity) ViewNumber() uint64 {

// 	countStr, _ := cache.RedisClient.Get(cache.ActivityViewKey(activity.ID)).Result()
// 	count, _ := strconv.ParseUint(countStr, 10, 64)
// 	return count
// }

//AddPaticiate 活动参与+1
func (activity *Activity) AddPaticiate() {
	cache.RedisClient.Incr(cache.ActivitPaticiteKey(activity.ID))
}

//PaticiateNumber 参与累计数
func (activity *Activity) PaticiateNumber() int {

	countStr, _ := cache.RedisClient.Get(cache.ActivitPaticiteKey(activity.ID)).Result()
	countUint64, _ := strconv.ParseUint(countStr, 10, 64)
	return int(countUint64)
}

// //AddWin 获得奖品数
// func (activity *Activity) AddWin() {
// 	cache.RedisClient.Incr(cache.ActivitWinKey(activity.ID))
// }

// //WinNumber 获奖累计数
// func (activity *Activity) WinNumber() uint64 {

// 	countStr, _ := cache.RedisClient.Get(cache.ActivitWinKey(activity.ID)).Result()
// 	count, _ := strconv.ParseUint(countStr, 10, 64)
// 	return count
// }
