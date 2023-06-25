package model

import (
	"github.com/jinzhu/gorm"
	"myBlog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 检查用户是否存在
func CheckUser(name string) (code int) {
	return 200
}

// 新增用户
func CreateUser(data *User) int {
	return errmsg.SUCCSE
}
