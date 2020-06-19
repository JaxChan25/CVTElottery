package api

import (
	"singo/service"
	"singo/util"

	"github.com/gin-gonic/gin"
)

// ActivityPost 新增活动
func ActivityPost(c *gin.Context) {

	service := service.ActivityPostService{}

	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

//ListActivities 列处所有活动
func ListActivities(c *gin.Context) {
	service := service.ListActivitiesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowActivity 取得单个活动详细页
func ShowActivity(c *gin.Context) {

	isMobile := util.IsMobile(c.Request.UserAgent())
	service := service.ShowActivityService{}
	res := service.Show(c.Param("id"), isMobile)
	c.JSON(200, res)
}

// UpdateActivity 更新某次活动
func UpdateActivity(c *gin.Context) {
	service := service.UpdateActivityService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetGraphData 取得传播数据
func GetGraphData(c *gin.Context) {

	service := service.GetGraphDataService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
