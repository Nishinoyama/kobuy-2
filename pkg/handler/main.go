package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/ent/user"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"net/http"
	"strconv"
	"time"
)

func GetUsersHandler(userClient *ent.UserClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		users, err := controller.GetAllUsers(userClient, cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, users)
	}
}

func FindUserHandler(userClient *ent.UserClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		userId, err := strconv.Atoi(gc.Param("id"))
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		user, err := controller.FindUser(userClient, cc, userId)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, user)
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

type PurchaseGroceryRequest struct {
	BuyerId   int `json:"buyer_id" binding:"required"`
	GroceryId int `json:"grocery_id" binding:"required"`
	Unit      int `json:"unit"`
}

func PurchaseGroceryHandler(client *ent.Client) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		var req PurchaseGroceryRequest
		if err := gc.BindJSON(&req); err != nil {
			gc.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := controller.PurchaseGrocery(client, cc, req.BuyerId, req.GroceryId, req.Unit); err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, true)
	}
}

func GetLedger(client *ent.Client) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		ledger, err := client.BalanceLog.Query().
			WithReceiver(func(query *ent.UserQuery) {
				query.Select(user.FieldName)
			}).
			WithDonor(func(query *ent.UserQuery) {
				query.Select(user.FieldName)
			}).All(cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, ledger)
	}
}
