package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowAddressService 查找地址的服务
type ShowAddressService struct {
}

// Show 展示地址
func (service *ShowAddressService) Show(id string) serializer.Response {

	var address model.Address

	err := model.DB.Where("game_user_id=?", id).First(&address).Error
	//err := model.DB.First(&address, id).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildAddressResponse(address),
	}
}
