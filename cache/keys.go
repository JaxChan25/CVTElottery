package cache

import (
	"fmt"
	"strconv"
)

const (

	//TotalRankKey 总排行
	TotalRankKey = "rank:total"
)

//SurplusTimesTotalActivityGameUserKey   surplustimes:total:活动:用户
func SurplusTimesTotalActivityGameUserKey(ActivityID uint, GamerUserID uint) string {
	return fmt.Sprintf("surplustimes:total:%s:%s", strconv.Itoa(int(ActivityID)), strconv.Itoa(int(GamerUserID)))
}

//SurplusTimesDailyActivityGameUserKey   surplustimes:daily:活动:用户
func SurplusTimesDailyActivityGameUserKey(ActivityID uint, GamerUserID uint) string {
	return fmt.Sprintf("surplustimes:daily:%s:%s", strconv.Itoa(int(ActivityID)), strconv.Itoa(int(GamerUserID)))
}

//redis中key的原则
//view:activity:1 -> 150
//表示id为1的activity的浏览量为150

//ActivityViewKey 生成活动浏览人数的Key
// func ActivityViewKey(id uint) string {
// 	return fmt.Sprintf("activity:view:%s", strconv.Itoa(int(id)))
// }

//ActivitPaticiteKey 生成抽奖参与人数的Key
func ActivitPaticiteKey(id uint) string {
	return fmt.Sprintf("activity:paticipate:%s", strconv.Itoa(int(id)))
}

// //ActivitWinKey 生成成功获奖人数的Key
// func ActivitWinKey(id uint) string {
// 	return fmt.Sprintf("activity:win:%s", strconv.Itoa(int(id)))
// }
