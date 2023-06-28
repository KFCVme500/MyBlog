package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "github.com/wejectchen/ginblog/api/v1"
	v2 "myBlog/api/v1"
	"myBlog/middleware"
	"myBlog/utils"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v2.GetUsers)
		auth.PUT("user/:id", v2.EditUser)
		auth.DELETE("user/:id", v2.DeleteUser)
		//修改密码
		///auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)
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
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)
		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
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
		router.POST("loginfront", v2.LoginFront)
		//
		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)
	}

	r.Run(utils.HttpPort)
}
