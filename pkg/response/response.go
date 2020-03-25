package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"visitor/pkg/ierr"
)

func ResponseOk(c *gin.Context, data interface{}) {
	var ret = map[string]interface{}{
		"code":    ierr.Success,
		"data":    data,
		"error":   "",
		"message": ierr.ErrorMessage[ierr.Success],
	}
	c.JSON(http.StatusOK, ret)
}

func ResponseErr(c *gin.Context, code int, err error) {
	var ret = map[string]interface{}{
		"code":    code,
		"data":    nil,
		"error":   "",
		"message": "",
	}
	if err != nil {
		ret["error"] = err.Error()
	}
	if message, ok := ierr.ErrorMessage[code]; ok {
		ret["message"] = message
	}
	c.JSON(http.StatusBadRequest, ret)
}
