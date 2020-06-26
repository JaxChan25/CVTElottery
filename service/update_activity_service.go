package service

import (
	"singo/model"
	"singo/serializer"

	"github.com/araddon/dateparse"
)

// UpdateActivityService 管理活动修改的服务
type UpdateActivityService struct {
	Title     string `json:"title"`
	LimitNum  int    `json:"limit_num"`  // '限制的抽奖次数',
	RuleText  string `json:"rule_text" ` // '活动规则介绍'
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// Post 用于修改活动
func (service *UpdateActivityService) Post(id string) serializer.Response {

	var activity model.Activity
	err := model.DB.First(&activity, id).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	t1, err := dateparse.ParseAny(service.StartTime)
	t2, err := dateparse.ParseAny(service.EndTime)

	err = model.DB.Model(&activity).Updates(model.Activity{
		Title:     service.Title,
		LimitNum:  service.LimitNum,
		RuleText:  service.RuleText,
		StartTime: t1,
		EndTime:   t2,
	}).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildActivityResponse(activity),
	}

}
