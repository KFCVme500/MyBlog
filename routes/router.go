package routes

import (
	"github.com/gin-gonic/gin"
	v2 "myBlog/api/v2"
	"myBlog/middleware"
	"myBlog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	auth := r.Group("api/v2")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v2.GetUsers)
		auth.PUT("user/:id", v2.EditUser)
		auth.DELETE("user/:id", v2.DeleteUser)
		//修改密码
		///auth.PUT("admin/changepw/:id", v2.ChangeUserPassword)
		// 分类模块的路由接口
		auth.GET("admin/category", v2.GetCate)
		auth.POST("category/add", v2.AddCategory)
		auth.PUT("category/:id", v2.EditCate)
		auth.DELETE("category/:id", v2.DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v2.GetArtInfo)
		auth.GET("admin/article", v2.GetArt)
		auth.POST("article/add", v2.AddArticle)
		auth.PUT("article/:id", v2.EditArt)
		auth.DELETE("article/:id", v2.DeleteArt)
		// 上传文件
		auth.POST("upload", v2.UpLoad)
		// 更新个人设置
		//auth.GET("admin/profile/:id", v2.GetProfile)
		//auth.PUT("profile/:id", v2.UpdateProfile)
		//// 评论模块
		//auth.GET("comment/list", v2.GetCommentList)
		//auth.DELETE("delcomment/:id", v2.DeleteComment)
		//auth.PUT("checkcomment/:id", v2.CheckComment)
		//auth.PUT("uncheckcomment/:id", v2.UncheckComment)
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

		// 文章分类信息模块
		router.GET("category", v2.GetCate)
		router.GET("category/:id", v2.GetCateInfo)

		// 文章模块
		router.GET("article", v2.GetArt)
		router.GET("article/list/:id", v2.GetCateArt)
		router.GET("article/info/:id", v2.GetArtInfo)

		// 登录控制模块
		router.POST("login", v2.Login)
		//router.POST("loginfront", v2.LoginFront)
		//
		//// 获取个人设置信息
		//router.GET("profile/:id", v2.GetProfile)
		//
		//// 评论模块
		//router.POST("addcomment", v2.AddComment)
		//router.GET("comment/info/:id", v2.GetComment)
		//router.GET("commentfront/:id", v2.GetCommentListFront)
		//router.GET("commentcount/:id", v2.GetCommentCount)
	}

	r.Run(utils.HttpPort)
}
