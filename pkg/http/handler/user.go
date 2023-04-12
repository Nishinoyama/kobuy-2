package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Client *ent.UserClient
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	cc := context.Background()
	users, err := controller.GetAllUsers(h.Client, cc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) Find(ctx *gin.Context) {
	cc := context.Background()
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	u, err := controller.FindUser(h.Client, cc, userId)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, err.Error())
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, u)
}
