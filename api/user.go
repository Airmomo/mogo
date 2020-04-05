package api

import (
	"mogo/cache"
	"mogo/model"
	"mogo/serializer"
	"mogo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(http.StatusOK, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	token, _ := c.Get("token")
	tokenString := token.(string)
	cache.RedisClient.SAdd("jwt:baned", tokenString)
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "登出成功",
	})
}
