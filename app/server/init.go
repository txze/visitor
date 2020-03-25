package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"visitor/app/server/router"
	"visitor/client/gorm_client"
)

func Init() *gin.Engine {
	var err error
	// 连接数据库
	_, err = gorm_client.Dial("mysql", viper.GetString("mysql.address"))
	if err != nil {
		panic(err)
	}

	// 初始化表
	err = gorm_client.Client.Master().Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1").
		AutoMigrate().Error
	if err != nil {
		panic(err)
	}

	// 初始化路由
	r := router.Init()
	return r
}
