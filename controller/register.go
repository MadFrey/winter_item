package controller

import (
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPost(c *gin.Context) {
	// 获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")
	if len(password) > 16 {
		util.PrintInfo(c, "密码长度超出限制！", 100)
		return
	} else if len(password) < 8 {
		util.PrintInfo(c, "密码太短！", 100)
		return
	}
	if password != rePassword {
		c.JSONP(http.StatusOK, gin.H{
			"code":    100,
			"message": "两次密码不一致",
		})
		return
	}
	flag := service.JudgeUserExist(username, password)
	if flag{
		c.JSONP(http.StatusOK,gin.H{
			"code":    100,
			"message": "用户名已经存在",
		})
		return
	}
	_,err:=service.AddNewUserProcess(username,password)
	if err != nil {
		c.JSONP(http.StatusOK,gin.H{
			"code":    805,
			"message": "注册失败",
		})
		return
	}
		c.JSONP(http.StatusOK,gin.H{
			"code":    0,
			"message": "注册成功！请登录",
		})
}
