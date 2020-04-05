package middleware

import (
	"github.com/gin-gonic/gin"
	"mogo/auth"
	"mogo/cache"
	"mogo/model"
	"mogo/serializer"
	"net/http"
	"strings"
)

//JWTAuthMiddleware 基于JWT的登录认证中间件
func JWTAuthMiddleware(head string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, serializer.Err(
				serializer.CodeLoseAuthHeader,
				"令牌不能为空，请登录获取令牌！",
				nil))
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != head {
			c.JSON(http.StatusOK, serializer.Err(
				serializer.CodeAuthWrong,
				"令牌格式有误！",
				nil))
			c.Abort()
			return
		}
		tokenString := parts[1]
		// parts[1]是获取到的tokenString，我们auth.user钟定义好的解析JWT的函数来解析它
		mc, err := auth.ParseUserToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, serializer.Err(
				serializer.CodeInvalidToken,
				"令牌已失效！",
				err))
			c.Abort()
			return
		}

		// 判断令牌是否在Redis黑名单里面，表示已注销
		if result, err := cache.RedisClient.SIsMember("jwt:baned", tokenString).Result(); result {
			c.JSON(http.StatusOK, serializer.Err(
				serializer.CodeInvalidToken,
				"令牌已注销！",
				err))
			c.Abort()
			return
		}

		// 将Token也放入Context, 用于用户登出添加黑名单
		c.Set("token", tokenString)

		user, err := model.GetUser(mc.UserID)
		if err == nil {
			// 将当前请求的user信息保存到请求的上下文c上
			c.Set("user", &user)
		} else {
			serializer.Err(
				serializer.CodeNeverLogin,
				"用户未登录或用户不存在！",
				err)
		}

		c.Next() // 后续的处理函数可以用过c.Get("user")来获取当前请求的用户信息
	}
}
