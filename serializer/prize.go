package serializer

import "singo/model"

// Prize 奖品序列化器
type Prize struct {
	ID         uint    `json:"id"`
	ActivityID uint    `json:"activity_id"` //外码
	Level      string  `json:"level"`       //'中奖等级',
	Name       string  `json:"name"`        //'奖品名',
	Prob       float64 `json:"prob"`        //中奖概率
	AllNum     int     `json:"all_num"`     //'总个数',
	SurplusNum int     `json:"surplus_num"` //剩余个数
	Image      string  `json:"image"`
	Ifwin      int     `json:"if_win"` //'抽中本项是否判定中奖'
	CreatedAt  string  `json:"created_at"`
}

// BuildArticle 序列化文章
func BuildArticle(prize model.GamePrize) Prize {

	return Prize{
		ID:         prize.ID,
		ActivityID: prize.ActivityID,
		Level:      prize.Level,
		Name:       prize.Name,
		Prob:       prize.Prob,
		AllNum:     prize.AllNum,
		SurplusNum: prize.SurplusNum,
		Image:      prize.Image,
		Ifwin:      prize.Ifwin,
		CreatedAt:  prize.CreatedAt.Format("2006-01-02 15:04"),
	}

}

// BuildPrizeResponse 序列化单个奖品响应
func BuildPrizeResponse(prize model.GamePrize) Prize {
	return BuildArticle(prize)
}

// BuildPrizesResponse 序列化多个奖品响应
func BuildPrizesResponse(items []model.GamePrize) (prizes []Prize) {

	for _, item := range items {
		prizes = append(prizes, BuildArticle(item))
	}
	return prizes
}
