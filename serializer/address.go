package serializer

import "singo/model"

// Address 地址序列化器
type Address struct {
	ID         uint   `json:"id" binding:"required"`
	GameUserID uint   `json:"game_user_id" binding:"required"` //外码
	Province   string `json:"province" binding:"required"`     //省
	City       string `json:"city" binding:"required"`         //市
	District   string `json:"district" binding:"required"`     //区
	Detail     string `json:"detail" binding:"required"`       //详细地址
	CreatedAt  string `json:"created_at"`
}

// BuildAddress 序列化地址
func BuildAddress(address model.Address) Address {

	return Address{
		ID:         address.ID,
		GameUserID: address.GameUserID,
		Province:   address.Province,
		City:       address.City,
		District:   address.District,
		Detail:     address.Detail,
		CreatedAt:  address.CreatedAt.Format("2006-01-02 15:04"),
	}

}

// BuildAddressResponse 序列化单个地址响应
func BuildAddressResponse(Address model.Address) Address {
	return BuildAddress(Address)
}

// // BuildPrizesResponse 序列化多个奖品响应
// func BuildPrizesResponse(items []model.GamePrize) (prizes []Prize) {

// 	for _, item := range items {
// 		prizes = append(prizes, BuildArticle(item))
// 	}
// 	return prizes
// }
