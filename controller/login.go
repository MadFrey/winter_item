package controller

import (
	"blog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)


func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := service.JudgeUserExist(username, password)
	if flag {
		s := sessions.Default(c)
		s.Set("loginUser", username)
		err := s.Save()
		if err != nil {
			panic(err)
		}
		tokenString,refreshTokenString:=service.CreateToken(username)
		c.JSON(http.StatusOK,gin.H{
			"code":    0,
			"message": "登录成功，欢迎进入！",
			"data":gin.H{
				"token":tokenString,
				"refreshToken":refreshTokenString,
			},
		})
	}else {
		c.JSONP(http.StatusOK,gin.H{
			"code":    1,
			"message": "用户名或密码错误！",
		})
	}

}

func LogoutGet(c *gin.Context)  {
	s:=sessions.Default(c)
	s.Delete("loginUser")
	_=s.Save()
	//重定向
	c.Redirect(http.StatusMovedPermanently,"https://blog.csdn.net/qq_36431213/article/details/82967982")
}
