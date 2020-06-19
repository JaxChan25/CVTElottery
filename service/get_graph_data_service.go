package service

import (
	"singo/model"
	"singo/serializer"
)

// GetGraphDataService 获得传播数据的服务
type GetGraphDataService struct {
	ActivityID uint   `json:"activity_id" binding:"required"`
	StartTime  string `json:"start_time" binding:"required"`
	EndTime    string `json:"end_time" binding:"required"`
}

// Get 获取传播数据
func (service *GetGraphDataService) Get() serializer.Response {

	/*
		浏览数据
	*/
	var viewGraphData []model.GraphData
	err := model.DB.Raw(
		`
	select concat(date_format(created_at,'%Y-%m-%d %H'),':',RPAD(floor(DATE_FORMAT(created_at,'%i')/5) * 5,2,0) ) as period, count(action_type) as count 
	from user_actions 
	where activity_id = ? and action_type =? and created_at between ? and ?
	group by period,action_type 
	`, service.ActivityID, 1, service.StartTime, service.EndTime).Scan(&viewGraphData).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	/*
		参与数据
	*/
	var paticipateGraphData []model.GraphData
	err = model.DB.Raw(
		`
	select concat(date_format(created_at,'%Y-%m-%d %H'),':',RPAD(floor(DATE_FORMAT(created_at,'%i')/5) * 5,2,0) ) as period, count(action_type) as count 
	from user_actions 
	where activity_id = ? and action_type =? and created_at between ? and ?
	group by period,action_type 
	`, service.ActivityID, 2, service.StartTime, service.EndTime).Scan(&paticipateGraphData).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	/*
		获奖数据
	*/
	var winGraphData []model.GraphData
	err = model.DB.Raw(
		`
	select concat(date_format(created_at,'%Y-%m-%d %H'),':',RPAD(floor(DATE_FORMAT(created_at,'%i')/5) * 5,2,0) ) as period, count(action_type) as count 
	from user_actions 
	where activity_id = ? and action_type =? and created_at between ? and ?
	group by period,action_type 
	`, service.ActivityID, 3, service.StartTime, service.EndTime).Scan(&winGraphData).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildAllGraphDataResponse(viewGraphData, paticipateGraphData, winGraphData),
		//Data: serializer.BuildOneGraphDataResponse(viewGraphData),
	}
}
