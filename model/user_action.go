package model

import "time"

// UserAction 模型
type UserAction struct {
	ID         uint `gorm:"primary_key"`
	GameUserID uint `gorm:"index"`                //外码
	ActivityID uint `gorm:"index"`                // 外码
	ActionType int  `gorm:"not null;default:'0'"` //'动作类型(1:浏览,2:参与抽奖,3:获奖)'
	Result     uint //'如果动作为抽奖或获奖，抽中的prize id',
	CreatedAt  time.Time
}
