package routes

import (
	"github.com/gin-gonic/gin"
	"myBlog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"msg": "ok",
			})
		})

	}
	r.Run(utils.HttpPort)
}
