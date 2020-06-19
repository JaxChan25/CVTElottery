package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// AddressPost 新增地址
func AddressPost(c *gin.Context) {

	service := service.AddressPostService{}

	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}


// ShowAddress 取得单个地址
func ShowAddress(c *gin.Context) {
	service := service.ShowAddressService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
