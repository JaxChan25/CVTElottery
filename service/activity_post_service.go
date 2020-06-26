package service

import (
	"singo/model"
	"singo/serializer"
	"singo/util"

	"github.com/araddon/dateparse"
)

// ActivityPostService 管理活动新增的服务
type ActivityPostService struct {
	GameManagerID   uint   `json:"game_manager_id" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Title           string `json:"title" binding:"required"`
	StartTime       string `json:"start_time"  binding:"required"`
	EndTime         string `json:"end_time"  binding:"required"`
	LimitType       int    `json:"limit_type" binding:"required"` //(0:无限制,1:每日抽奖次数限制,2:总抽奖次数限制)',
	LimitNum        int    `json:"limit_num" binding:"required"`  // '限制的抽奖次数',
	RuleText        string `json:"rule_text" binding:"required"`  // '活动规则介绍'
	BannerImage     string `json:"banner_image"`
	LotteryImage    string `json:"lottery_image"`
	BackgrouldColor string `json:"background_color"`
	VirtualNum      int    `json:"virtual_num" binding:"required"` //'虚拟参与者个数'
}

// Post 用于新建活动
func (service *ActivityPostService) Post() serializer.Response {

	t1, err := dateparse.ParseAny(service.StartTime)
	t2, err := dateparse.ParseAny(service.EndTime)
	if err != nil {
		return serializer.ParamErr("时间解析失败", err)
	}

	activity := model.Activity{
		GameManagerID:   service.GameManagerID,
		Name:            service.Name,
		Title:           service.Title,
		StartTime:       t1,
		EndTime:         t2,
		LimitType:       service.LimitType,
		LimitNum:        service.LimitNum,
		RuleText:        service.RuleText,
		BannerImage:     service.BannerImage,
		LotteryImage:    service.LotteryImage,
		BackgrouldColor: service.BackgrouldColor,
		VirtualNum:      service.VirtualNum,
	}

	util.Log().Info(activity.StartTime.Format("2006-01-02 15:04" + "\n"))
	util.Log().Info(activity.EndTime.Format("2006-01-02 15:04" + "\n"))

	if err := model.DB.Create(&activity).Error; err != nil {
		return serializer.ParamErr("活动写入失败", err)
	}

	return serializer.Response{
		Data: serializer.BuildActivityResponse(activity),
	}

}
