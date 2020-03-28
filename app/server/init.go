package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"visitor/app/server/router"
	"visitor/client/gorm_client"
	"visitor/client/redis_client"
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

	//初始化redis
	redis_client.MasterPool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "redis.master.address")
		},
	}

	redis_client.SalvePool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "redis.salve.address")
		},
	}
	// 初始化路由
	r := router.Init()
	return r
}
