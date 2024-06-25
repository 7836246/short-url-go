package initialize

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"short-url-go/internal/global"
	"short-url-go/internal/model"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(global.G_CONFIG.Http.DbName), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	// 创建 short_url 表
	err = db.AutoMigrate(&model.ShortUrl{})
	if err != nil {
		fmt.Println("自动建表失败")
		panic(err)
	}
	return db
}
