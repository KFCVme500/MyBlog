package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wejectchen/ginblog/api/v1"
	"github.com/wejectchen/ginblog/middleware"
	"myBlog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.GET("admin/user", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口

		//文章模块的路由接口

	}
	r.Run(utils.HttpPort)
}
