package middleware

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "请先登录",
			})
			c.Abort()
			return
		}
		parts := strings.Split(authorization, " ")
		if !(len(parts) == 3 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "请先登录",
			})
			c.Abort()
			return
		}
		Token, _, err := service.ParseToken(parts[1], parts[2])
		if err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "请重新登录",
			})
			c.Abort()
			return
		}
		c.Set("username", Token.Username)
		c.Next()
	}
}
