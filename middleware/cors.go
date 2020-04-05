package middleware

import (
	"mogo/util"
	"os"
	"regexp"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie","Access-Control-Allow-Origin"}
	config.ExposeHeaders= []string{"Content-Length", "Access-Control-Allow-Origin"}

	// 设置GIN的环境
	// 默认环境设置为debug,这里通过判断环境变量GIN_MODE是否为release切换到生产环境
	mode := os.Getenv("GIN_MODE")
	if strings.EqualFold(mode, "release") {
		gin.SetMode(os.Getenv("GIN_MODE"))
	}
	// 输出GIN的环境
	util.Log().Info("gin-mode:%s", gin.Mode())

	// 运行在Release模式下会进行跨域保护
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境需要配置跨域域名，否则403
		// string数组可实现多跨域
		config.AllowOrigins = []string{"http://www.example.com"}
	} else {
		// 测试环境下模糊匹配本地开头的请求
		config.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
		//支持所有请求
		//config.AllowAllOrigins = true
	}
	config.AllowCredentials = true
	return cors.New(config)
}
