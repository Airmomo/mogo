package model

//执行数据迁移

func migration() {
	// 自动迁移模式(mysql自动创建结构体的database)
	// Set防止自动创建数据库时使用编码错误
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
}
