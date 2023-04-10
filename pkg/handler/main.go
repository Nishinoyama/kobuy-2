package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"net/http"
	"time"
)

func GetUsersHandler(userClient *ent.UserClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		users, err := userClient.Query().All(cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, users)
	}
}

func GetGroceriesHandler(groceryClient *ent.GroceryClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		users, err := groceryClient.Query().WithProvider().All(cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, users)
	}
}

type ProvideGroceryRequest struct {
	ProviderId     int        `json:"provider_id,omitempty" binding:"required"`
	Name           string     `json:"name,omitempty" binding:"required"`
	Price          int        `json:"price,omitempty" binding:"required"`
	Unit           int        `json:"unit,omitempty" binding:"required"`
	ExpirationDate *time.Time `json:"expiration_date"`
}

func ProvideGroceryHandler(groceryClient *ent.GroceryClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		var req ProvideGroceryRequest
		if err := gc.BindJSON(&req); err != nil {
			gc.JSON(http.StatusBadRequest, err.Error())
			return
		}
		g := groceryClient.Create().
			SetProviderID(req.ProviderId).
			SetName(req.Name).
			SetPrice(req.Price).
			SetUnit(req.Unit)
		if req.ExpirationDate != nil {
			g.SetExpirationDate(*req.ExpirationDate)
		}
		if grocery, err := g.Save(cc); err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			gc.JSON(http.StatusOK, grocery)
			return
		}
	}
}
