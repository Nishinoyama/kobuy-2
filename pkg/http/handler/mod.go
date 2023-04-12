package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/ent/ledger"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"net/http"
)

type PurchaseGroceryRequest struct {
	BuyerId   int `json:"buyer_id" binding:"required"`
	GroceryId int `json:"grocery_id" binding:"required"`
	Unit      int `json:"unit"`
}

func NewPurchaseGroceryHandler(r *gin.RouterGroup, client *ent.Client) {
	r.POST("/purchase", func(gc *gin.Context) {
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
	})
}

type CashLedgerRequest struct {
	DonorId    int `json:"donor_id" binding:"required"`
	ReceiverId int `json:"receiver_id" binding:"required"`
	Price      int `json:"price" binding:"required"`
}

func NewCashHandler(r *gin.RouterGroup, client *ent.Client) {
	r.POST("/cash", func(gc *gin.Context) {
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
	})
}
