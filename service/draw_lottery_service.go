package service

import (
	"math/rand"
	"singo/cache"
	"singo/model"
	"singo/serializer"
	"singo/util"
	"strconv"
	"time"
	"unsafe"

	"github.com/go-redis/redis"
)

// DrawLotteryService 某个用户对某个活动进行抽奖的服务
type DrawLotteryService struct {
	ActivityID uint `json:"activity_id" binding:"required"`
	GameUserID uint `json:"game_user_id" binding:"required"`
}

// Post  进行抽奖
func (service *DrawLotteryService) Post() serializer.Response {

	var activity model.Activity
	err := model.DB.Preload("GamePrizes").First(&activity, service.ActivityID).Error
	if err != nil {
		return serializer.ParamErr("活动查询失败", err)
	}

	// util.Log().Info(time.Now().Format("2006-01-02 15:04" + "\n"))
	// util.Log().Info(activity.StartTime.Format("2006-01-02 15:04" + "\n"))
	// util.Log().Info(activity.EndTime.Format("2006-01-02 15:04" + "\n"))

	if time.Now().Before(activity.StartTime) || time.Now().After(activity.EndTime) {
		return serializer.ParamErr("活动尚未开始或者活动已经结束", err)
	}

	var key string
	if activity.LimitType == 1 { //每日
		key = cache.SurplusTimesDailyActivityGameUserKey(service.ActivityID, service.GameUserID)
	} else { //总共
		key = cache.SurplusTimesTotalActivityGameUserKey(service.ActivityID, service.GameUserID)
	}

	val, err := cache.RedisClient.Get(key).Result()

	surplusTimes64, _ := strconv.ParseInt(val, 10, 64)
	surplusTimes := *(*int)(unsafe.Pointer(&surplusTimes64))

	if err == redis.Nil { //key不存在，则需要设置redis

		if activity.LimitType == 1 { //每日
			cache.RedisClient.Set(key, strconv.Itoa(activity.LimitNum), 24*60*60*time.Second)
			surplusTimes = activity.LimitNum
		} else { //总共
			cache.RedisClient.Set(key, strconv.Itoa(activity.LimitNum), 0)
			surplusTimes = activity.LimitNum
		}
	}

	//剩余抽奖次数大于0才可以进行抽奖
	if surplusTimes > 0 {

		//抽奖行为记录
		cache.RedisClient.Decr(key)
		activity.AddPaticiate()
		userAction := model.UserAction{
			GameUserID: service.GameUserID,
			ActivityID: service.ActivityID,
			ActionType: 2, //'动作类型(1:浏览,2:参与抽奖,3:获奖)'
		}
		if err := model.DB.Create(&userAction).Error; err != nil {
			return serializer.ParamErr("浏览记录写入失败", err)
		}

		/*
			执行抽奖算法
		*/
		var prizes []model.GamePrize = activity.GamePrizes
		var index int = DrawPrizeIndex(prizes)
		if prizes[index].Ifwin == 2 { //抽中

			//记录抽中
			userAction = model.UserAction{
				GameUserID: service.GameUserID,
				ActivityID: service.ActivityID,
				ActionType: 3,
				Result:     prizes[index].ID,
			}
			if err := model.DB.Create(&userAction).Error; err != nil {
				return serializer.ParamErr("浏览记录写入失败", err)
			}

			prizes[index].SurplusNum--
			model.DB.Save(&prizes[index])
		}

		return serializer.Response{
			Data: serializer.BuildPrizeResponse(prizes[index]),
		}

	}

	return serializer.Response{
		Data: "次数少不能抽奖",
	}
}

//DrawPrizeIndex 进行抽奖，返回中奖的index; 注意是index,不是prize_key
func DrawPrizeIndex(prizes []model.GamePrize) int {

	var drawRandom float64 = rand.Float64()

	//得先把剩余奖品为0的奖品去掉
	for i := 0; i < len(prizes); i++ {
		if prizes[i].SurplusNum == 0 {
			drawRandom -= prizes[i].Prob
		}

	}

	for i := 0; i < len(prizes); i++ {
		if prizes[i].SurplusNum == 0 {
			continue
		}

		if drawRandom-prizes[i].Prob < 0 {
			return i
		} else {
			drawRandom -= prizes[i].Prob
		}
	}

	util.Log().Error("抽奖算法出错了！")
	//事实上不会从这里进行返回
	return 0
}
