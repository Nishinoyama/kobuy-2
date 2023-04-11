package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/ent/ledger"
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
		u, err := controller.FindUser(userClient, cc, userId)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, u)
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
		l, err := client.Ledger.Query().
			WithReceiver(func(query *ent.UserQuery) {
				query.Select(user.FieldName)
			}).
			WithPayer(func(query *ent.UserQuery) {
				query.Select(user.FieldName)
			}).All(cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, l)
	}
}

type CashLedgerRequest struct {
	DonorId    int `json:"donor_id" binding:"required"`
	ReceiverId int `json:"receiver_id" binding:"required"`
	Price      int `json:"price" binding:"required"`
}

func CashLedgerHandler(client *ent.Client) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		var req CashLedgerRequest
		err := gc.BindJSON(&req)
		if err != nil {
			gc.JSON(http.StatusBadRequest, err.Error())
			return
		}

		tx, err := client.Tx(cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if err := tx.User.UpdateOneID(req.DonorId).AddBalance(-req.Price).Exec(cc); err != nil {
			if err := tx.Rollback(); err != nil {
				gc.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if err := tx.User.UpdateOneID(req.ReceiverId).AddBalance(req.Price).Exec(cc); err != nil {
			if err := tx.Rollback(); err != nil {
				gc.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		l, err := tx.Ledger.Create().
			SetPayerID(req.DonorId).
			SetReceiverID(req.ReceiverId).
			SetPrice(req.Price).
			SetType(ledger.TypeCash).Save(cc)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				gc.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if err := tx.Commit(); err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		gc.JSON(http.StatusOK, l)
	}
}
