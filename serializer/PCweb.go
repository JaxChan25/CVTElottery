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

type uintData struct {
	ViewCount        int `json:"view_count" binding:"required"`
	ParticipateCount int `json:"participate_count" binding:"required"`
	WinCount         int `json:"win_count" binding:"required"`
}

// OneRecondOfGraphData 图表的一条信息
type OneRecondOfGraphData struct {
	Period   string   `json:"period" binding:"required"`
	UintData uintData `json:"uint_data" binding:"required"`
}

//BuildOneRecordGraphDataResponse 序列化一条数据
func BuildOneRecordGraphDataResponse(record model.GraphData) OneRecondOfGraphData {

	return OneRecondOfGraphData{
		Period: record.Period,
		UintData: uintData{
			ViewCount:        record.UintData.ViewCount,
			ParticipateCount: record.UintData.ParticipateCount,
			WinCount:         record.UintData.WinCount,
		},
	}
}

// BuildGraphDataResponse 序列表格数据（包含多条数据）
func BuildGraphDataResponse(items []model.GraphData) (datagraph []OneRecondOfGraphData) {
	for _, item := range items {
		datagraph = append(datagraph, BuildOneRecordGraphDataResponse(item))
	}
	return datagraph
}
