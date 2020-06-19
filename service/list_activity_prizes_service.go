package service

import (
	"singo/model"
	"singo/serializer"
)

// ListActivityPrizesService 处理某个活动的中奖记录的服务
type ListActivityPrizesService struct {
	ActivityID uint `json:"activity_id" binding:"required"`
	Limit      int  `form:"limit"`
	Offset     int  `form:"offset"`
}

// List 处理某个活动的中奖记录
func (service *ListActivityPrizesService) List() serializer.Response {

	var results []model.ListActivityPrizesResult

	if service.Limit == 0 {
		service.Limit = 6
	}

	err := model.DB.Raw(
		`
	select  t4.user_name ,t4.real_name ,t4.mobile,t1.created_at,t2.name as prize_name,t2.level as prize_level,t3.province ,t3.city,t3.district ,t3.detail 
	from ((((select  created_at, game_user_id ,activity_id,result as prize_id  from user_actions where action_type =? and activity_id =?) as t1
	inner join game_prizes t2 on t2.activity_id = t1.activity_id and t2.id = t1.prize_id))
	inner join addresses t3 on t3.game_user_id = t1.game_user_id)
	inner join game_users t4 on t4.id = t1.game_user_id
	limit ? offset ?
	`, 3, service.ActivityID, service.Limit, service.Offset).Scan(&results).Error

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}
	total := len(results)
	return serializer.BuildListResponse(serializer.BuildListActivityPrizesResultsResponse(results), uint(total))

}
