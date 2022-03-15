package controller

import (
	"blog/service"
	"blog/util"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
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
		tokenString, refreshTokenString := service.CreateToken(username)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "登录成功，欢迎进入！",
			"data": gin.H{
				"token":        tokenString,
				"refreshToken": refreshTokenString,
			},
		})
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"code":    1,
			"message": "用户名或密码错误！",
		})
	}

}

func LogoutGet(c *gin.Context) {
	s := sessions.Default(c)
	s.Delete("loginUser")
	_ = s.Save()
	//重定向到主页
	c.Redirect(http.StatusMovedPermanently, "https://blog.csdn.net/qq_36431213/article/details/82967982")
}

func HTMLGet(c *gin.Context) {
	c.HTML(http.StatusOK, "html.html", gin.H{
		"title": "html",
	})
}

func Oauth(c *gin.Context) {
	code := c.Query("code")
	var err error
	tokenUrl := util.GetTokenAuthUrl(code)
	var token *util.Token
	if token, err = util.GetToken(tokenUrl); err != nil {
		log.Println(err)
		return
	}
	userInfo, err := util.GetUserInfo(token)
	if err != nil {
		log.Println(err)
		return
	}
	id := int(userInfo["id"].(float64))
	err = service.JudgeUserWithGitId(id)
	if err != nil {
		avatarUrl := userInfo["avatar_url"].(string)
		nickname := userInfo["name"].(string)
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		userCode := fmt.Sprintf("%08v", rnd.Int31n(1000000000))
		username := userCode
		_, err = service.AddNewUserProcess(username, "123456")
		if err != nil {
			util.PrintInfo(c, "注册失败", 1)
			log.Println(err)
			return
		}
		err := service.UpdateInfoWithGithub(username, avatarUrl, nickname, id)
		if err != nil {
			util.PrintInfo(c, "更新数据失败", 1)
			return
		}
		tokenString, refreshTokenString := service.CreateToken(username)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "登录成功,您的用户名为" + username + "初始密码为123456",
			"data": gin.H{
				"token":        tokenString,
				"refreshToken": refreshTokenString,
			},
		})
	} else {
		username, err := service.QueryUserWithGitId(id)
		if err != nil {
			log.Println(err)
			util.PrintInfo(c, "查询失败", 1)
			return
		}
		tokenString, refreshTokenString := service.CreateToken(username)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "登录成功，欢迎进入！",
			"data": gin.H{
				"token":        tokenString,
				"refreshToken": refreshTokenString,
			},
		})
	}
}
