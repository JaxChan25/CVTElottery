package model

// GraphData 数据表格
type GraphData struct {
	Period   string
	UintData UintData
}

//UintData 浏览、参与、中奖数据
type UintData struct {
	ViewCount        int
	ParticipateCount int
	WinCount         int
}
