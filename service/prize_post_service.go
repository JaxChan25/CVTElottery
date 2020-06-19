package service

import (
	"singo/model"
	"singo/serializer"
)

// PrizePostService 管理奖品新增的服务
type PrizePostService struct {
	ActivityID uint    `json:"activity_id" binding:"required"` //外码
	Level      string  `json:"level" binding:"required"`       //'中奖等级',
	Name       string  `json:"name" binding:"required"`        //'奖品名',
	Prob       float64 `json:"prob" binding:"required"`        //中奖概率
	AllNum     int     `json:"all_num" binding:"required"`     //'总个数',
	SurplusNum int     `json:"surplus_num" binding:"required"` //剩余个数
	Image      string  `json:"image"`
	Ifwin      int     `json:"if_win" binding:"required"` //'抽中本项是否判定中奖'
}

// Post 用于新建奖品
func (service *PrizePostService) Post() serializer.Response {

	prize := model.GamePrize{
		ActivityID: service.ActivityID,
		Level:      service.Level,
		Name:       service.Name,
		Prob:       service.Prob,
		AllNum:     service.AllNum,
		SurplusNum: service.SurplusNum,
		Image:      service.Image,
		Ifwin:      service.Ifwin,
	}

	if err := model.DB.Create(&prize).Error; err != nil {
		return serializer.ParamErr("奖品写入失败", err)
	}

	return serializer.Response{
		Data: serializer.BuildPrizeResponse(prize),
	}
}
