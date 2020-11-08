package handler

import (
	"github.com/gin-gonic/gin"
	"time"
)

func GetGoodsInfo(c *gin.Context) {
	// Parse JSON
	var json struct {
		Sign string `json:"sign" binding:"required"`
		BizData string `json:"biz_data" binding:"required"`
		Test string `json:"test" binding:"required"`
	}

	if err := c.Bind(&json); err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"message": "get good info",
		"data": json,
	})
}

func GetGoodsSkus(c *gin.Context) {
	time.Sleep(time.Second * 5)
	c.JSON(200, gin.H{
		"code": 0,
		"message": "get good sku list",
		"data": nil,
	})
}