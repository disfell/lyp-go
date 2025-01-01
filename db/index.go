package db

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	//var err error
	//DB, err = gorm.Open(sqlite.Open("lyp.db"), &gorm.Config{})
	//if err != nil {
	//	logger.Errorf("创建sqlite连接失败: %+v", err)
	//}
	//
	//// 自动迁移模式
	//DB.AutoMigrate(&model.Article{})
}
