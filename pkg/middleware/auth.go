package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserKey = "user"

func AuthMiddleware(authIsRequired bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get(UserKey)
		if userId == nil {
			if authIsRequired {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
