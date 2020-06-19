package service

import (
	"singo/cache"
	"singo/model"
	"singo/serializer"
	"strconv"
	"time"
	"unsafe"

	"github.com/go-redis/redis"
)

// GetSurplusTimesService 剩余抽奖次数的服务
type GetSurplusTimesService struct {
	ActivityID uint `json:"activity_id" binding:"required"`
	GameUserID uint `json:"game_user_id" binding:"required"`
}

// Get 获取剩余抽奖次数
func (service *GetSurplusTimesService) Get() serializer.Response {

	var activity model.Activity
	err := model.DB.First(&activity, service.ActivityID).Error
	if err != nil {
		return serializer.ParamErr("活动查询失败", err)
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

	return serializer.Response{
		Data: surplusTimes,
	}
}
