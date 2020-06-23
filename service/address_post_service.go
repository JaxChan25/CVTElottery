package service

import (
	"singo/model"
	"singo/serializer"
)

// AddressPostService 管理地址新增的服务
type AddressPostService struct {
	GameUserID uint   `json:"game_user_id" binding:"required"`
	RealName   string `form:"real_name" json:"real_name" binding:"required,min=1,max=4"`
	Mobile     string `form:"mobile" json:"mobile" binding:"required,min=11,max=11"`
	Province   string `json:"province" binding:"required"` //省
	City       string `json:"city" binding:"required"`     //市
	District   string `json:"district" binding:"required"` //区
	Detail     string `json:"detail" binding:"required"`   //详细地址
}

// Post 用于新建地址
func (service *AddressPostService) Post() serializer.Response {

	address := model.Address{
		GameUserID: service.GameUserID,
		Province:   service.Province,
		RealName:   service.RealName,
		Mobile:     service.Mobile,
		City:       service.City,
		District:   service.District,
		Detail:     service.Detail,
	}

	if err := model.DB.Create(&address).Error; err != nil {
		return serializer.ParamErr("地址写入失败", err)
	}

	return serializer.Response{
		Data: serializer.BuildAddressResponse(address),
	}
}
