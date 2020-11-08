package handler

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"message": "get user info",
		"data": nil,
	})
}

func GetUserVip(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"message": "get user vip",
		"data": nil,
	})
}