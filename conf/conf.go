package conf

import (
	"mogo/cache"
	"mogo/model"
	"mogo/util"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 设置Jwt.StandardClaims的基本配置
 	util.BuildJwtStandardClaims(time.Hour*2,os.Getenv("JWT_SECRET"),os.Getenv("JWT_ISSUER"),os.Getenv("JWT_SigningMethod"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接MySQL数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// 连接Redis缓存
	cache.Redis()
}
