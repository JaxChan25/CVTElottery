package serializer

import "singo/model"

// ListActivityPrizesResult 查询活动中奖记录的序列化器
type ListActivityPrizesResult struct {
	UserName   string `json:"user_name" binding:"required"`
	RealName   string `json:"real_name" binding:"required"`
	Mobile     string `json:"mobile" binding:"required"`
	CreateAt   string `json:"create_at" binding:"required"`
	PrizeName  string `json:"prize_name" binding:"required"`
	PrizeLevel string `json:"prize_level" binding:"required"`
	Province   string `json:"province" binding:"required"`
	City       string `json:"city" binding:"required"`
	District   string `json:"district" binding:"required"`
	Detail     string `json:"detail" binding:"required"`
}

// BuildListActivityPrizesResult 序列化活动中奖记录
func BuildListActivityPrizesResult(result model.ListActivityPrizesResult) ListActivityPrizesResult {

	return ListActivityPrizesResult{
		UserName:   result.UserName,
		RealName:   result.RealName,
		Mobile:     result.Mobile,
		CreateAt:   result.CreatedAt.Format("2006-01-02 15:04"),
		PrizeName:  result.PrizeName,
		PrizeLevel: result.PrizeLevel,
		Province:   result.Province,
		City:       result.City,
		District:   result.District,
		Detail:     result.Detail,
	}

}

// BuildListActivityPrizesResultResponse 序列化单个
func BuildListActivityPrizesResultResponse(result model.ListActivityPrizesResult) ListActivityPrizesResult {
	return BuildListActivityPrizesResult(result)
}

// BuildListActivityPrizesResultsResponse 序列化多个
func BuildListActivityPrizesResultsResponse(items []model.ListActivityPrizesResult) (results []ListActivityPrizesResult) {

	for _, item := range items {
		results = append(results, BuildListActivityPrizesResultResponse(item))
	}
	return results
}

/*
	图表
*/

// OneRecondOfGraphData 图表的一条信息
type OneRecondOfGraphData struct {
	Period string `json:"period" binding:"required"`
	Count  int    `json:"count" binding:"required"`
}

// AllGraphData 图表的所有信息
type AllGraphData struct {
	View        DataList `json:"view" binding:"required"`
	Participate DataList `json:"participate" binding:"required"`
	Win         DataList `json:"win" binding:"required"`
}

//BuildOneRecordGraphDataResponse 序列化一条数据
func BuildOneRecordGraphDataResponse(record model.GraphData) OneRecondOfGraphData {
	return OneRecondOfGraphData{
		Period: record.Period,
		Count:  record.Count,
	}
}

// BuildOneGraphDataResponse 序列化单个表格（包含多条数据）
func BuildOneGraphDataResponse(items []model.GraphData) (datagraph []OneRecondOfGraphData) {
	for _, item := range items {
		datagraph = append(datagraph, BuildOneRecordGraphDataResponse(item))
	}
	return datagraph
}

// BuildAllGraphDataResponse 序列化三个表格（包含多条数据）
func BuildAllGraphDataResponse(viewGraphData []model.GraphData, paticipateGraphData []model.GraphData, winGraphData []model.GraphData) AllGraphData {

	return AllGraphData{
		View:        DataList{Items: BuildOneGraphDataResponse(viewGraphData), Total: uint(len(viewGraphData))},
		Participate: DataList{Items: BuildOneGraphDataResponse(paticipateGraphData), Total: uint(len(viewGraphData))},
		Win:         DataList{Items: BuildOneGraphDataResponse(winGraphData), Total: uint(len(viewGraphData))},
	}
}
