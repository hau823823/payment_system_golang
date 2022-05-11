package rsp

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//正確狀態
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

//錯誤狀態
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"error": msg,
		})
}
