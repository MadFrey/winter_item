package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func RefreshMyTokenGet(c *gin.Context)  {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(200, gin.H{
			"code": 2003,
			"msg":  "请求头中auth为空",
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSON(200, gin.H{
			"code": 2004,
			"msg":  "请求头中auth格式有误",
		})
		return
	}
	parseToken, isUpd, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSON(200, gin.H{
			"code": 2005,
			"msg":  "无效的Token",
		})
		return
	}
	if isUpd {
		parts[1], parts[2] = service.CreateToken(parseToken.Username)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "成功",
			"data": gin.H{
				"accessToken":  parts[1],
				"refreshToken": parts[2],
			},
		})
	}
	c.Set("username", parseToken.Username)
}
