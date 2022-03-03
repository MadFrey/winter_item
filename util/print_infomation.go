package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PrintInfo(c *gin.Context,str string,status int)  {
		c.JSON(http.StatusOK,gin.H{
			"message":str,
			"status":status,
		})
}