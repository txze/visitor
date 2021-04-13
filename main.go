package main

import (
	_ "visitor/config"

	_ "visitor/logger"

	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"time"
	"visitor/app/server"
)

func main() {

	// 设置随机数因子
	rand.Seed(time.Now().Unix())

	// 启动服务
	port := viper.GetString("server.port")
	log.Info("Server listening on ", port)

	router := server.Init()
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println(err)
		log.Info("Server Stop, Error: ", err)
	}
}
