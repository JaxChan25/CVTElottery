package service

import (
	"singo/model"
	"singo/serializer"
	"time"

	"github.com/araddon/dateparse"
)

// GetGraphDataService 获得传播数据的服务
type GetGraphDataService struct {
	ActivityID uint   `json:"activity_id" binding:"required"`
	StartTime  string `json:"start_time" binding:"required"`
	EndTime    string `json:"end_time" binding:"required"`
}

//Result mysql返回的结果
type Result struct {
	Period string
	Count  int
}

// Get 获取传播数据
func (service *GetGraphDataService) Get() serializer.Response {

	/*
		浏览数据
	*/
	var graphData []model.GraphData

	/*
		生成一个map作为容器，预生成数据
	*/
	t1, err := dateparse.ParseAny(service.StartTime)
	t2, err := dateparse.ParseAny(service.EndTime)

	for {

		str := t1.Format("2006-01-02 15:04")
		//util.Log().Info(str + "\n")
		graphData = append(graphData,
			model.GraphData{
				Period: str,
				UintData: model.UintData{
					ViewCount:        0,
					ParticipateCount: 0,
					WinCount:         0,
				},
			},
		)
		t1 = t1.Add(time.Hour * 1)
		if t1.After(t2) {
			break
		}
	}

	/*
		浏览人数
	*/
	var results []Result
	err = model.DB.Raw(
		`
	SELECT DATE_FORMAT(created_at, '%Y-%m-%d %H:00') AS period, COUNT(*) AS count
	FROM user_actions ua 
	WHERE activity_id = ? and action_type =? and created_at between ? and ? 
	GROUP BY period
	ORDER BY period
	`, service.ActivityID, 1, service.StartTime, service.EndTime).Scan(&results).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	j := 0
	for i := 0; i < len(results); i++ {
		//util.Log().Info(results[i].Period + "\n")
		for {

			if results[i].Period != graphData[j].Period {
				j++
			} else {
				graphData[j].UintData.ViewCount = results[i].Count
				break
			}
		}
	}
	j = 0
	results = results[0:0]

	/*
		抽奖人数
	*/
	err = model.DB.Raw(
		`
	SELECT DATE_FORMAT(created_at, '%Y-%m-%d %H:00') AS period, COUNT(*) AS count
	FROM user_actions ua 
	WHERE activity_id = ? and action_type =? and created_at between ? and ? 
	GROUP BY period
	ORDER BY period
	`, service.ActivityID, 2, service.StartTime, service.EndTime).Scan(&results).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	for i := 0; i < len(results); i++ {
		//util.Log().Info(results[i].Period + "\n")
		for {

			if results[i].Period != graphData[j].Period {
				j++
			} else {
				graphData[j].UintData.ParticipateCount = results[i].Count
				break
			}
		}
	}
	j = 0
	results = results[0:0]

	/*
		中奖人数
	*/
	err = model.DB.Raw(
		`
	SELECT DATE_FORMAT(created_at, '%Y-%m-%d %H:00') AS period, COUNT(*) AS count
	FROM user_actions ua 
	WHERE activity_id = ? and action_type =? and created_at between ? and ? 
	GROUP BY period
	ORDER BY period
	`, service.ActivityID, 3, service.StartTime, service.EndTime).Scan(&results).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	for i := 0; i < len(results); i++ {
		//util.Log().Info(results[i].Period + "\n")
		for {

			if results[i].Period != graphData[j].Period {
				j++
			} else {
				graphData[j].UintData.WinCount = results[i].Count
				break
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildGraphDataResponse(graphData),
		//Data: serializer.BuildOneGraphDataResponse(viewGraphData),
	}
}
