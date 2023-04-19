package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Controller *controller.UserController
}

func NewUserHandler(r *gin.RouterGroup, uc *controller.UserController) {
	handler := UserHandler{Controller: uc}

	user := r.Group("/users")
	{
		user.GET("/", handler.GetAll)
		user.GET("/:id", handler.Find)
	}
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	res, err := h.Controller.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *UserHandler) Find(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Controller.Find(userId)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, err.Error())
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, res)
}
