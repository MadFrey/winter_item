package controller

import (
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
)

func VCodeGet(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		util.PrintInfo(c, "未登录或请求过期", 110)
	}

	u := service.GetUserinfo(username.(string))
	err, code := util.SendEmail(u.Email, "验证码")
	if err != nil {
		util.PrintInfo(c, "发送验证码失败", 806)
		return
	}
	util.StoreVerCode(code)
}
