package service

import (
	"singo/model"
	"singo/serializer"
	"strconv"
)

// ShowActivityService 活动详情的服务
type ShowActivityService struct {
}

// Show 展示活动
func (service *ShowActivityService) Show(id string, isMobile bool) serializer.Response {

	var activity model.Activity
	err := model.DB.Preload("GamePrizes").First(&activity, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	idUint, _ := strconv.ParseUint(id, 10, 64)
	if isMobile {
		userAction := model.UserAction{
			ActivityID: uint(idUint),
			ActionType: 1, //'动作类型(1:浏览,2:参与抽奖,3:获奖)'
		}

		if err := model.DB.Create(&userAction).Error; err != nil {
			return serializer.ParamErr("浏览记录写入失败", err)
		}
	}

	return serializer.Response{
		Data: serializer.BuildActivityResponse(activity),
	}
}
