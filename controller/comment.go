package controller

import (
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

func AllCommentGet(c *gin.Context) {
	Model, _ := strconv.Atoi(c.Query("Model"))
	targetId, _ := strconv.Atoi(c.Query("target_id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		if Model == 2 {
			commentList := service.NoTokenSecondQueryComment(targetId, page, size)
			c.JSONP(http.StatusOK, gin.H{
				"message": "success",
				"code":    1,
				"data":    commentList,
			})
			return
		}

		commentList := service.NoTokenQueryComment(targetId, page, size)
		c.JSONP(http.StatusOK, gin.H{
			"message": "success",
			"code":    1,
			"data":    commentList,
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		if Model == 2 {
			commentList := service.NoTokenSecondQueryComment(targetId, page, size)
			c.JSONP(http.StatusOK, gin.H{
				"message": "success",
				"code":    1,
				"data":    commentList,
			})
			return
		}

		commentList := service.NoTokenQueryComment(targetId, page, size)
		c.JSONP(http.StatusOK, gin.H{
			"message": "success",
			"code":    1,
			"data":    commentList,
		})
		return
	}
	Token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		if Model == 2 {
			commentList := service.NoTokenSecondQueryComment(targetId, page, size)
			c.JSONP(http.StatusOK, gin.H{
				"message": "success",
				"code":    1,
				"data":    commentList,
			})
			return
		}
		commentList := service.NoTokenQueryComment(targetId, page, size)
		c.JSONP(http.StatusOK, gin.H{
			"message": "success",
			"code":    1,
			"data":    commentList,
		})
		return
	}
	if Model == 2 {
		CommentList:=service.QuerySecondComment(targetId, page, size, Token.Username)
		c.JSONP(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data":    CommentList,
		})
		return
	}
	CommentList := service.QueryComment(targetId, page, size,  Token.Username)
	c.JSONP(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    CommentList,
	})
	return
}

func CommentPost(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSONP(http.StatusOK, gin.H{
			"message": "?????????",
			"code":    0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK, gin.H{
			"message": "?????????",
			"code":    0,
		})
		return
	}
	Token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":    0,
			"message": "?????????????????????????????????",
		})
		return
	}
	Model, err := strconv.Atoi(c.PostForm("model"))
	if err != nil {
		log.Println(err)
		return
	}
	TargetId, err := strconv.Atoi(c.PostForm("TargetId"))
	if err != nil {
		log.Println(err)
		return
	}
	content := c.PostForm("content")
	files := c.Request.MultipartForm.File["file"]
	var images []string
	for _, file := range files {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "????????????!?????????png,jpg,gif,jpeg??????",
			})
			return
		}
		newFileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(file.Filename)
		err = c.SaveUploadedFile(file, "./files/"+newFileName)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "?????????????????????",
			})
			return
		}
		ip := "127.0.0.1:9090"
		url := ip + "/static/" + newFileName
		images = append(images, url)
	}
	imageString := util.ArrayToString(images)
	commentId, err := service.PostComment(TargetId, content, imageString, Model, Token.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "??????????????????",
			"code":    0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    1,
		"data":    commentId,
	})
}

func UpdateCommentPut(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSONP(http.StatusOK, gin.H{
			"message": "?????????",
			"code":    0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK, gin.H{
			"message": "?????????",
			"code":    0,
		})
		return
	}
	_, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":    0,
			"message": "?????????????????????????????????",
		})
		return
	}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	content := c.PostForm("content")
	files := c.Request.MultipartForm.File["file"]
	var images []string
	for _, file := range files {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "????????????!?????????png,jpg,gif,jpeg??????",
			})
			return
		}
		newFileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(file.Filename)
		err = c.SaveUploadedFile(file, "./files/"+newFileName)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "?????????????????????",
			})
			return
		}
		ip := "127.0.0.1:9090"
		url := ip + "/static/" + newFileName
		images = append(images, url)
	}
	imageString := util.ArrayToString(images)
	_, err = service.UpdateComment(commentId, content, imageString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "??????????????????",
			"code":    0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "??????????????????",
		"code":    1,
	})
}

func CommentDelete(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSONP(http.StatusOK, gin.H{
			"message": "?????????",
			"code":    0,
		})
		return
	}
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		c.JSONP(http.StatusOK, gin.H{
			"message": "?????????",
			"code":    0,
		})
		return
	}
	Token, _, err := service.ParseToken(parts[1], parts[2])
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"code":    0,
			"message": "?????????????????????????????????",
		})
		return
	}
	CommentId, err := strconv.Atoi(c.Param("CommentId"))
	if err != nil {
		log.Println(err)
		return
	}
	_, ok:= service.DeleteComment(CommentId,Token.Username)
	if ok==false {
		c.JSON(http.StatusOK, gin.H{
			"message": "??????????????????",
			"code":    1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "??????????????????",
		"code":    0,
	})
}
