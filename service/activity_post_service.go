package service

import (
	"singo/model"
	"singo/serializer"
	"time"
)

// ActivityPostService 管理活动新增的服务
type ActivityPostService struct {
	GameManagerID   uint      `json:"game_manager_id" binding:"required"`
	Name            string    `json:"name" binding:"required"`
	Title           string    `json:"title" binding:"required"`
	Type            int       `json:"type"`                     //活动种类，0：大转盘
	State           int       `json:"state" binding:"required"` //活动状态(0:未开始,1:已开始,-1:已结束)
	Mode            int       `json:"mode" binding:"required"`  //'活动模式(0:系统活动,1:独立活动) 本项目中总为1',
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	LimitType       int       `json:"limit_type" binding:"required"` //(0:无限制,1:每日抽奖次数限制,2:总抽奖次数限制)',
	LimitNum        int       `json:"limit_num" binding:"required"`  // '限制的抽奖次数',
	RuleText        string    `json:"rule_text" binding:"required"`  // '活动规则介绍'
	BannerImage     string    `json:"banner_image"`
	LotteryImage    string    `json:"lottery_image"`
	BackgrouldColor string    `json:"backgrould_color"`
	VirtualNum      int       `json:"virtual_num" binding:"required"` //'虚拟参与者个数'
}

// Post 用于新建活动
func (service *ActivityPostService) Post() serializer.Response {

	activity := model.Activity{
		GameManagerID:   service.GameManagerID,
		Name:            service.Name,
		Title:           service.Title,
		Type:            service.Type,
		State:           service.State,
		Mode:            service.Mode,
		StartTime:       time.Now(),
		EndTime:         time.Now(),
		LimitType:       service.LimitType,
		LimitNum:        service.LimitNum,
		RuleText:        service.RuleText,
		BannerImage:     service.BannerImage,
		LotteryImage:    service.LotteryImage,
		BackgrouldColor: service.BackgrouldColor,
		VirtualNum:      service.VirtualNum,
	}

	if err := model.DB.Create(&activity).Error; err != nil {
		return serializer.ParamErr("活动写入失败", err)
	}

	return serializer.Response{
		Data: serializer.BuildActivityResponse(activity),
	}

}
