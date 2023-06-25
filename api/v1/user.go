package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/utils/validator"
	"myBlog/model"
	"myBlog/utils/errmsg"
	"net/http"
)

//查询用户是否存在

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	- = c.ShouldBindJSON(&data)
	//校验器，
	msg,validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCSE {
		c.JSON(
				http.StatusOK,gin.H{
					"status": validCode,
					"message": msg,
			},
			)
		c.Abort()
		return
	}
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}

	c.JSON(
			http.StatusOK,gin.H{
				"status":code,
				"message":errmsg.GetErrMsg(code),
		},
		)
}
//查询单个用户

//查询用户列表

//编辑用户
//删除用户
