package router

import (
	"blog/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)
var store = cookie.NewStore([]byte("secret"))

func CreateRouter()*gin.Engine {
	r:=gin.Default()
	r.StaticFS("/static", http.Dir("./files"))
	r.Use(sessions.Sessions("sessionId", store))
	r.POST("/user/token", controller.LoginPost)
	r.GET("/logout", controller.LogoutGet)
	r.POST("/user/register",controller.RegisterPost)
	r.PUT("/user/info",controller.ChangUserinfoPost)
	r.GET("/user/token/refresh",controller.RefreshMyTokenGet)
	r.GET("/user/info/:user_id",controller.UserinfoGet)
	r.PUT("/user/password",controller.ChangeUserPwdPut)
	postgres:=r.Group("/post")
	{
		postgres.GET("/list",controller.ArticleListGet)
		postgres.GET("/single/:PostId",controller.PostContentGet)
		postgres.POST("/single",controller.PostSinglePost)
		postgres.PUT("/single/:PostId",controller.UpdatePostPut)
		postgres.GET("/search",controller.ArticleSearchGet)
		postgres.DELETE("/single/:PostId",controller.ArticleDelete)
	}
	r.GET("/topic/list",controller.AllTopicGet)

	r.GET("/comment",controller.AllCommentGet)
	r.POST("/comment",controller.CommentPost)
	r.PUT("/comment/:commentId",controller.UpdateCommentPut)
	r.DELETE("/comment/:CommentId",controller.CommentDelete)

	 v:= r.Group("/operate")
	{
		v.PUT("/praise",controller.PraisedPut)
		v.PUT("/focus",controller.FocusUserPut)
		v.GET("/focus/list",controller.FocusListGet)
		v.PUT("/collect",controller.CollectPut)
		v.GET("/collect/list",controller.CollectListGet)
	}
	return r
}
