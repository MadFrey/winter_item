package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

func ChangUserinfoPost(c *gin.Context)  {
	authorization:= c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "更改用户信息失败",
		})
		return
	}
	parts := strings.Split(authorization, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "更改用户信息失败",
		})
		return
	}
	Token,_, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改失败，请重新登录",
		})
		return
	}
	username := Token.Username
	user:=service.GetUserinfo(username)
	email := c.DefaultPostForm("email",user.Email)
	NickName := c.DefaultPostForm("NickName",user.NickName)
	Gender := c.DefaultPostForm("Gender",user.Gender)
	qq,_:= strconv.Atoi(c.DefaultPostForm("qq",strconv.Itoa(user.Qq)))
	introduction := c.DefaultPostForm("introduction",user.Introduction)
	birthday:= c.DefaultPostForm("birthday",user.Birthday)
	phone,_:= strconv.Atoi(c.DefaultPostForm("phone",strconv.Itoa(user.Phone)))
	file,err := c.FormFile("AvatarUrl")
	if err != nil {
		_,err=service.UpdatePersonalInfoNoAvatar(username,email,NickName,Gender,introduction,qq,birthday,phone)
		if err != nil {
			log.Println(err)
			c.JSONP(http.StatusOK,gin.H{
				"code": 0,
				"msg":  "修改用户信息失败！",
			})
			return
		}
		c.JSONP(http.StatusOK,gin.H{
			"message":"修改成功",
			"code": 1,
		})
		return
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
		})
		return
	}
	newFileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(file.Filename)
	err = c.SaveUploadedFile(file, "./files/"+newFileName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "头像上传失败！",
		})
		return
	}
	ip := "127.0.0.1:9090"
	AvatarUrl := ip + "/static/" + newFileName
	_,err=service.UpdatePersonalInfo(username,email,NickName,AvatarUrl,Gender,introduction,qq,birthday,phone)
	if err != nil {
		c.JSONP(http.StatusOK,gin.H{
			"message":"修改信息失败",
			"code": 0,
		})
	}
	c.JSONP(http.StatusOK,gin.H{
		"message":"修改成功",
		"code": 1,
	})
}
