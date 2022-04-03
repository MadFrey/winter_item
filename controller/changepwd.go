package controller

import (
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ChangeUserPwdPut(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	parts := strings.Split(authorization, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	parseToken, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请重新登录",
		})
		return
	}
	oldPassword := c.PostForm("oldPassword")
	VCode := c.PostForm("vCode")
	ok := service.JudgeUserExist(parseToken.Username, oldPassword)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "密码错误",
		})
		return
	}
	ok = util.VerifyCode(VCode)
	if !ok {
		util.PrintInfo(c, "验证码错误", 1)
	} else {
		newPassword := c.PostForm("newPassword")
		_, err = service.UpdateUserPwd(parseToken.Username, newPassword)
		if err != nil {
			panic(err)
		}
		c.JSONP(http.StatusOK, gin.H{
			"code":    0,
			"message": "密码修改成功",
		})
	}
}
