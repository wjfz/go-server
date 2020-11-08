package routers

import (
	"github.com/gin-gonic/gin"
	"wserver/handler"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"message": "hello world",
		"data": nil,
	})
}

func InitRouter(r *gin.Engine) {
	r.GET("/", helloHandler)

	userRouters := r.Group("/user")
	userRouters.GET("/info", handler.GetUserInfo)
	userRouters.GET("/vip", handler.GetUserVip)

	goodsRouters := r.Group("/goods")
	goodsRouters.POST("/info", handler.GetGoodsInfo)
	goodsRouters.GET("/skus", handler.GetGoodsSkus)
}
