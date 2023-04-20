package handler

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"github.com/nishinoyama/kobuy-2/pkg/ent/user"
	"github.com/nishinoyama/kobuy-2/pkg/middleware"
	"net/http"
)

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func LoginHandler(r *gin.Engine, uc *ent.UserClient) {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	r.POST("/login", login(uc))
	r.OPTIONS("/login", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})
	r.GET("/logout", logout())

	private := r.Group("/authed")
	private.Use(middleware.AuthMiddleware(true))
	{
		private.GET("/me", func(c *gin.Context) {
			userId := c.GetInt("user_id")
			c.JSON(http.StatusOK, gin.H{"user_id": userId})
		})
	}
}

func login(uc *ent.UserClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		u, err := uc.Query().Where(user.NameEQ(req.UserName)).First(context.TODO())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusBadRequest, "Incorrect")
				return
			} else {
				c.JSON(http.StatusInternalServerError, err)
				return
			}
		}
		session := sessions.Default(c)
		session.Set(middleware.UserKey, u.ID)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": u.ID})
	}
}

func logout() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, "logout")
	}
}
