package controller

import (
	"blog/model"
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

func ArticleListGet(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	if authHeader == "" {
		listContent := service.NoTokenMakeArticleList(pageNum, size)
		c.JSONP(http.StatusOK, gin.H{
			"code":    1,
			"message": "success",
			"data":    listContent,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		listContent := service.NoTokenMakeArticleList(pageNum, size)
		c.JSONP(http.StatusOK, gin.H{
			"code":    1,
			"message": "success",
			"data":    listContent,
		})
		return
	}
	Token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		listContent := service.NoTokenMakeArticleList(pageNum, size)
		c.JSONP(http.StatusOK, gin.H{
			"code":    1,
			"message": "success",
			"data":    listContent,
		})
		return
	}
	listContent := service.MakeArticleList(pageNum, size, Token.Username)
	c.JSONP(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    listContent,
	})
	return
}

func PostContentGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("PostId"))
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		article := service.NoTokenQuerySingleArticleProcess(id)
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "success",
			"data":    article,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		article := service.NoTokenQuerySingleArticleProcess(id)
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "success",
			"data":    article,
		})
		return
	}
	Token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		article := service.NoTokenQuerySingleArticleProcess(id)
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "success",
			"data":    article,
		})
		return
	}
	article, err := service.QuerySingleArticleProcess(id, Token.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "查询失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    article,
	})
}

func PostSinglePost(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	topicId := c.PostForm("topicId")
	files := c.Request.MultipartForm.File["file"]
	var images []string
	for _, file := range files {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			})
			return
		}
		newFileName := strconv.FormatInt(time.Now().Unix(),10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(file.Filename)
		err = c.SaveUploadedFile(file, "./files/"+newFileName)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "文件上传失败！",
			})
			return
		}
		ip:="127.0.0.1:9090"
		url:=ip+"/static/"+newFileName
		images = append(images, url)
	}
	imageString := util.ArrayToString(images)
	u1 := service.GetUserinfo(token.Username)
	var article model.Article
	article.TopicId = topicId
	article.Title = title
	article.Content = content
	article.Pictures = imageString
	article.Author = u1.NickName
	article.Username = token.Username
	article.Avatar = u1.AvatarUrl
	article.CreateTime = time.Now()
	PostId, err := service.AddArticleProcess(article)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"PostId":  PostId,
	})
}

func UpdatePostPut(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	Token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	PostId, err := strconv.Atoi(c.Param("PostId"))
	if err != nil {
		log.Fatal(err)
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	topicId := c.PostForm("topicId")
	files := c.Request.MultipartForm.File["file"]
	article, err := service.QuerySingleArticleProcess(PostId, Token.Username)
	if err != nil {
		log.Println(err)
		return
	}
	var images []string
	for _, file := range files {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
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
				"msg":  "文件上传失败！",
			})
			return
		}
		ip := "127.0.0.1:9090"
		url := ip + "/static/" + newFileName
		images = append(images, url)
	}
	article.TopicId = topicId
	article.Title = title
	article.Content = content
	article.Id = PostId
	_, err = service.UpdateArticleProcess(article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "更改文章失败",
			"code":    0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"info": "success",
		"code": 1,
	})
}

func ArticleSearchGet(c *gin.Context) {
	key := c.Query("key")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Println(err)
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		log.Println(err)
	}
	searchList := service.SearchArticlesProcess(key, page, size)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    1,
		"data":    searchList,
	})
}

func AllTopicGet(c *gin.Context) {
	topicList := service.QueryAllTopicsProcess()
	c.JSONP(http.StatusOK, gin.H{
		"message": "success",
		"code":    1,
		"data":    topicList,
	})
}

func ArticleDelete(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	Token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请先登录",
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("PostId"))
	_, ok:= service.DeleteArticleProcess(id,Token.Username)
	if ok ==false {
		c.JSONP(http.StatusOK, gin.H{
			"message": "删除失败",
			"code":    0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    1,
	})
}
