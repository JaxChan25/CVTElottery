package serializer

import (
	"singo/model"
)

// Activity 活动序列化器
type Activity struct {
	ID              uint    `json:"id"`
	GameManagerID   uint    `json:"game_manager_id" binding:"required"`
	GamePrizes      []Prize `json:"game_prizes" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	Title           string  `json:"title" binding:"required"`
	StartTime       string  `json:"start_time"`
	EndTime         string  `json:"end_time"`
	LimitType       int     `json:"limit_type" binding:"required"` //(0:无限制,1:每日抽奖次数限制,2:总抽奖次数限制)',
	LimitNum        int     `json:"limit_num" binding:"required"`  // '限制的抽奖次数',
	RuleText        string  `json:"rule_text" binding:"required"`  // '活动规则介绍'
	BannerImage     string  `json:"banner_image" binding:"required"`
	LotteryImage    string  `json:"lottery_image" binding:"required"`
	BackgrouldColor string  `json:"background_color" binding:"required"`
	VirtualNum      int     `json:"virtual_num" binding:"required"` //'虚拟参与者个数'
	CreatedAt       string  `json:"created_at"`
	ParticipateNum  int     `json:"participate_num" binding:"required"` //真实参与个数
}

// BuildActivity 序列化活动
func BuildActivity(activity model.Activity) Activity {

	return Activity{
		ID:              activity.ID,
		GamePrizes:      BuildPrizesResponse(activity.GamePrizes),
		GameManagerID:   activity.GameManagerID,
		Name:            activity.Name,
		Title:           activity.Title,
		StartTime:       activity.StartTime.Format("2006-01-02 15:04"),
		EndTime:         activity.EndTime.Format("2006-01-02 15:04"),
		LimitType:       activity.LimitType,
		LimitNum:        activity.LimitNum,
		RuleText:        activity.RuleText,
		BannerImage:     activity.BannerImage,
		LotteryImage:    activity.LotteryImage,
		BackgrouldColor: activity.BackgrouldColor,
		VirtualNum:      activity.VirtualNum,
		CreatedAt:       activity.CreatedAt.Format("2006-01-02 15:04"),
		ParticipateNum:  activity.PaticiateNumber(),
	}

}

// BuildActivityResponse 序列化单个活动响应
func BuildActivityResponse(activity model.Activity) Activity {
	return BuildActivity(activity)
}

// BuildActivitiesResponse 序列化多个活动响应
func BuildActivitiesResponse(items []model.Activity) (activities []Activity) {

	for _, item := range items {
		activities = append(activities, BuildActivity(item))
	}
	return activities
}
