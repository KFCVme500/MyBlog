package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/middleware"
	v2 "myBlog/api/v2"
	"myBlog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	auth := r.Group("api/v2")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.GET("admin/user")
		//auth.PUT("user/:id", v2.EditUser)
		//auth.DELETE("user/:id", v2.DeleteUser)
		//分类模块的路由接口

		//文章模块的路由接口

	}
	/*
		前端展示页面接口
	*/
	router := r.Group("api/v2")
	{
		// 用户信息模块
		router.POST("user/add", v2.AddUser)
		router.GET("user/:id", v2.GetUserInfo)
		router.GET("users", v2.GetUsers)
	}
	r.Run(utils.HttpPort)
}
