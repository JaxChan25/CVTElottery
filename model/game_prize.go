package model

import (
	"github.com/jinzhu/gorm"
)

// GamePrize 奖品模型
type GamePrize struct {
	gorm.Model
	ActivityID uint    `gorm:"NOT NULL;INDEX"`            //外码
	Level      string  `gorm:"type:varchar(20);not null"` //'中奖等级',
	Name       string  `gorm:"type:varchar(20);not null"` //'奖品名',
	Prob       float64 `gorm:"not null"`                  //中奖概率
	AllNum     int     `gorm:"not null;default:0"`        //'总个数',
	SurplusNum int     `gorm:"not null;default:0"`        //剩余个数
	Image      string  `gorm:"type:varchar(150);not null;default:'PrizeImage_default.png'"`
	Ifwin      int     `gorm:"not null;default:0"` //'抽中本项是否判定中奖' 0为不中奖
}
