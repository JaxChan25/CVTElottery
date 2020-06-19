package service

import (
	"singo/model"
	"singo/serializer"
)

// ListActivitiesService 活动列表的服务
type ListActivitiesService struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// List 展示奖品列表
func (service *ListActivitiesService) List() serializer.Response {
	var activies []model.Activity
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	//文章总数
	if err := model.DB.Model(model.Activity{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Offset).Find(&activies).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildActivitiesResponse(activies), uint(total))
}
