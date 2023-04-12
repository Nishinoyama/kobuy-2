package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"net/http"
	"time"
)

type GroceryHandler struct {
	Client *ent.GroceryClient
}

func (h *GroceryHandler) GetAll(ctx *gin.Context) {
	cc := context.Background()
	users, err := h.Client.Query().WithProvider().All(cc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, users)
}

type ProvideGroceryRequest struct {
	ProviderId     int        `json:"provider_id,omitempty" binding:"required"`
	Name           string     `json:"name,omitempty" binding:"required"`
	Price          int        `json:"price,omitempty" binding:"required"`
	Unit           int        `json:"unit,omitempty" binding:"required"`
	ExpirationDate *time.Time `json:"expiration_date"`
}

func (h *GroceryHandler) Provide(ctx *gin.Context) {
	cc := context.Background()
	var req ProvideGroceryRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	g := h.Client.Create().
		SetProviderID(req.ProviderId).
		SetName(req.Name).
		SetPrice(req.Price).
		SetUnit(req.Unit)
	if req.ExpirationDate != nil {
		g.SetExpirationDate(*req.ExpirationDate)
	}
	if grocery, err := g.Save(cc); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, grocery)
		return
	}
}
