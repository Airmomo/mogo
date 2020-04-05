package service

import (
	"github.com/gin-gonic/gin"
	"mogo/auth"
	"mogo/model"
	"mogo/serializer"
	"mogo/util"
	"os"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 生成tokenString
	tokenString, err := auth.GetUserToken(user.ID)
	if err != nil {
		return serializer.Err(
			serializer.CodeUnCreateToken,
			"令牌获取失败",
			err)
	}

	return serializer.Response{
		Data: gin.H{
			"access_token": tokenString,
			"expires_in":   util.GetExpiresTime(),
			"token_type":   os.Getenv("JWT_HEAD"),
		},
	}
}
