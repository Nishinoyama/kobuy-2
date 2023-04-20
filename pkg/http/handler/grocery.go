package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"github.com/nishinoyama/kobuy-2/pkg/middleware"
	"net/http"
	"strconv"
)

type GroceryHandler struct {
	Controller *controller.GroceryController
}

func NewGroceryHandler(r *gin.RouterGroup, gc *controller.GroceryController) {
	handler := GroceryHandler{gc}

	grocery := r.Group("/groceries")
	{
		grocery.GET("/", handler.GetAll)
		grocery.GET("/:id", handler.Find)

		authed := grocery.Group("")
		authed.Use(middleware.AuthRequiredMiddleware())
		{
			authed.POST("/provide", handler.Provide)
			authed.OPTIONS("/provide", func(ctx *gin.Context) {
				ctx.JSON(http.StatusNoContent, true)
			})
		}
	}
}

func (h *GroceryHandler) GetAll(ctx *gin.Context) {
	res, err := h.Controller.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *GroceryHandler) Find(ctx *gin.Context) {
	groceryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Controller.Find(groceryId)
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

func (h *GroceryHandler) Provide(ctx *gin.Context) {
	var req controller.GroceryProvideRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Controller.Provide(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}
