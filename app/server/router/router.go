package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	return r
}
