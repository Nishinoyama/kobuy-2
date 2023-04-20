package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserIdKey = "user_id"
const UserNameKey = "user_name"

func AuthMiddleware(authIsRequired bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get(UserIdKey)
		userName := session.Get(UserNameKey)
		c.Set("authed", userName != nil)
		if userName == nil {
			if authIsRequired {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
			return
		}
		c.Set("user_id", userId)
		c.Set("user_name", userName)
		c.Next()
	}
}
