package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// PrizePost 新增奖品
func PrizePost(c *gin.Context) {

	service := service.PrizePostService{}

	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// ListPrizes 列处所有奖品
func ListPrizes(c *gin.Context) {
	service := service.ListPrizesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListUserPrizes 列出某用户在某活动中奖的所有奖品
func ListUserPrizes(c *gin.Context) {
	service := service.ListUserPrizesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListActivityPrizes 列出某活动的中奖记录
func ListActivityPrizes(c *gin.Context) {
	service := service.ListActivityPrizesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
