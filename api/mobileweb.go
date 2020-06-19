package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// GetSurplusTime 取得某个用户在某个活动的剩余次数
func GetSurplusTime(c *gin.Context) {

	service := service.GetSurplusTimesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DrawLottery 取得某个用户在某个活动的抽奖结果
func DrawLottery(c *gin.Context) {

	service := service.DrawLotteryService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
