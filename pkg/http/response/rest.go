package response

import (
	"time"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/ginting/pkg/utils/helper"
)


func DefaultAppHeader(c *gin.Context){
	tm 		:= time.Now().Unix()
	c.Writer.Header().Set("App-Name", 	 helper.MustGetEnv("APP_NAME"))
	c.Writer.Header().Set("App-Version", helper.MustGetEnv("APP_VERSION"))
	c.Writer.Header().Set("App-Time", 	 strconv.Itoa(int(tm)))
}

func Json(c *gin.Context, data interface{}) {
	DefaultAppHeader(c)
	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"data":          data,
		"error_message": nil,
	})
	return
}

func NoContent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}

func Accepted(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusAccepted)
}