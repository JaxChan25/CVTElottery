package service

import (
	"singo/model"
	"singo/serializer"
)

// ListPrizesService 奖品列表的服务
type ListPrizesService struct {
}

// List 展示奖品列表
func (service *ListPrizesService) List() serializer.Response {

	var prizes []model.GamePrize
	total := 0

	if err := model.DB.Model(model.GamePrize{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Find(&prizes).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildPrizesResponse(prizes), uint(total))

}
