package controller

import (
	"context"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"github.com/nishinoyama/kobuy-2/pkg/service"
)

type PurchaseGroceryRequest struct {
	BuyerId   int `json:"buyer_id" binding:"required"`
	GroceryId int `json:"grocery_id" binding:"required"`
	Unit      int `json:"unit" binding:"required"`
}

func Purchase(client *ent.Client, req PurchaseGroceryRequest) error {
	return service.Purchase(context.TODO(), client, req.BuyerId, req.GroceryId, req.Unit)
}

type CashLedgerRequest struct {
	DonorId    int `json:"donor_id" binding:"required"`
	ReceiverId int `json:"receiver_id" binding:"required"`
	Price      int `json:"price" binding:"required"`
}

func Cash(client *ent.Client, req CashLedgerRequest) (*ent.Ledger, error) {
	return service.Cash(context.TODO(), client, req.DonorId, req.ReceiverId, req.Price)
}
