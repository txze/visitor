package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"visitor/model"
)

func ServerTokenCheck(c *gin.Context) {
	var token = getToken(c)
	var err error
	if token == "" {
		err = errors.New("invalid token")
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// 查询token的有效值
	uid, err := CheckToken(token)
	if err != nil {
		logrus.WithFields(logrus.Fields{"token": token}).WithError(err).Error("ServerTokenCheck check token error")
		err = errors.New("check token panic")
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//
	c.Set("user_id", uid)
	c.Next()
}

func getToken(c *gin.Context) string {
	var token = c.GetHeader("token")
	if token != "" {
		return token
	}
	if c.Request.Method == "POST" || c.Request.Method == "PUT" {
		token = c.PostForm("token")
	}
	if token != "" {
		return token
	}
	return c.Query("token")
}
func CheckToken(token string) (uid int, err error) {
	userModel := &model.User{Token: token}
	_, err = userModel.LoadByToken()
	if err != nil {
		return
	}
	uid = userModel.Id
	return
}
