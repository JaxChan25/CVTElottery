package service

import (
	"singo/model"
	"singo/serializer"
)

// UpdateActivityService 管理活动修改的服务
type UpdateActivityService struct {
	Title    string `json:"title"`
	LimitNum int    `json:"limit_num"`  // '限制的抽奖次数',
	RuleText string `json:"rule_text" ` // '活动规则介绍'

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
	err = model.DB.Model(&activity).Updates(model.Activity{
		Title:    service.Title,
		LimitNum: service.LimitNum,
		RuleText: service.RuleText,
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
