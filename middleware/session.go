package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func UseSession() gin.HandlerFunc {
	store:=cookie.NewStore([]byte("loginUser"))
	return sessions.Sessions("loginUser",store)
}
