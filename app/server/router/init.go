package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	// 游戏服务端路由注册
	r = Router(r)
	return r
}
