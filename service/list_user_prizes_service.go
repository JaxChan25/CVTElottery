package service

import (
	"singo/model"
	"singo/serializer"
)

// ListUserPrizesService 某个用户中奖奖品列表的服务
type ListUserPrizesService struct {
	ActivityID uint `json:"activity_id" binding:"required"`
	GameUserID uint `json:"game_user_id" binding:"required"`
}

// List 某个用户中奖奖品列表的服务
func (service *ListUserPrizesService) List() serializer.Response {

	var prizes []model.GamePrize

	err := model.DB.Raw(
		`SELECT * FROM game_prizes where id IN (SELECT result FROM user_actions WHERE game_user_id=? and activity_id=? and action_type=?)`,service.GameUserID, service.ActivityID,3).Scan(&prizes).Error

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}
	total := len(prizes)
	return serializer.BuildListResponse(serializer.BuildPrizesResponse(prizes), uint(total))

}
