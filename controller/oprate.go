package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func PraisedPut(c *gin.Context)  {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == ""{
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	Token,_,err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":0,
			"message":"登录已过期，请重新登录",
		})
		return
	}
	TargetId,_:=strconv.Atoi(c.PostForm("TargetId"))
	Model,_:=strconv.Atoi(c.PostForm("Model"))
	err=service.IsPraiseProcess(TargetId,Token.Username,Model)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message":"点赞失败",
			"code":0,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		"code":1,
	})
}

func FocusUserPut(c *gin.Context)  {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == ""{
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	Token,_,err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":0,
			"message":"登录已过期，请重新登录",
		})
		return
	}
	userId:=c.PostForm("user_id")
	ok:=service.FocusUser(Token.Username,userId)
	if ok == false {
		c.JSON(http.StatusOK,gin.H{
			"message":"关注失败",
			"code":0,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		"code":1,
	})
}

func FocusListGet(c *gin.Context)  {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == ""{
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	Token,_,err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":0,
			"message":"登录已过期，请重新登录",
		})
		return
	}
	list:=service.QueryFocusList(Token.Username)
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		"code":1,
		"data":list,
	})
}

func CollectPut(c *gin.Context)  {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == ""{
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	Token,_,err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":0,
			"message":"登录已过期，请重新登录",
		})
		return
	}
	postId,_:=strconv.Atoi(c.PostForm("post_id"))
	ok:=service.CollectProcess(postId,Token.Username)
	if  ok==false {
		c.JSONP(http.StatusOK,gin.H{
			"code":0,
			"message":"收藏失败",
		})
		return
	}
	c.JSONP(http.StatusOK,gin.H{
		"code":1,
		"message":"success",
	})
}

func CollectListGet(c *gin.Context)  {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == ""{
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK,gin.H{
			"message":"未登录",
			"code":0,
		})
		return
	}
	Token,_,err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":0,
			"message":"登录已过期，请重新登录",
		})
		return
	}
	collectList:=service.QueryCollectList(Token.Username)
	c.JSONP(http.StatusOK,gin.H{
		"code":1,
		"message":"success",
		"data":collectList,
	})
}

