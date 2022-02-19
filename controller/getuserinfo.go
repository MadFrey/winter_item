package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserinfoGet(c *gin.Context) {
	username:=c.Param("user_id")
	u:=service.GetUserinfo(username)
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"avatar":u.AvatarUrl,
		"nickname":u.NickName,
		"introduction":u.Introduction,
		"phone":u.Phone,
		"qq":u.Qq,
		"gender":u.Gender,
		"email":u.Email,
		"birthday":u.Birthday,
	})
}