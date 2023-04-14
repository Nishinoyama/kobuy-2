package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"net/http"
)

func NewPurchaseHandler(r *gin.RouterGroup, client *ent.Client) {
	r.POST("/purchase", func(gc *gin.Context) {
		var req controller.PurchaseGroceryRequest
		if err := gc.BindJSON(&req); err != nil {
			gc.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := controller.Purchase(client, req); err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, true)
	})
}

func NewCashHandler(r *gin.RouterGroup, client *ent.Client) {
	r.POST("/cash", func(gc *gin.Context) {
		var req controller.CashLedgerRequest
		err := gc.BindJSON(&req)
		if err != nil {
			gc.JSON(http.StatusBadRequest, err.Error())
			return
		}

		l, err := controller.Cash(client, req)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		gc.JSON(http.StatusOK, l)
	})
}
