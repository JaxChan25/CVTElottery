package service

import (
	"singo/model"
	"singo/serializer"
)

// ListPrizesService 奖品列表的服务
type ListPrizesService struct {
}

// List 展示奖品列表
func (service *ListPrizesService) List(id string) serializer.Response {

	var prizes []model.GamePrize

	if err := model.DB.Where("activity_id=?", id).Find(&prizes).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildPrizesResponse(prizes), uint(len(prizes)))

}
