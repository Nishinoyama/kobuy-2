package controller

import (
	"context"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/pkg/service"
	"time"
)

type GroceryController struct {
	GroceryService *service.GroceryService
}

type GroceryProvideRequest struct {
	ProviderId     int        `json:"provider_id,omitempty" binding:"required"`
	Name           string     `json:"name,omitempty" binding:"required"`
	Price          int        `json:"price,omitempty" binding:"required"`
	Unit           int        `json:"unit,omitempty" binding:"required"`
	ExpirationDate *time.Time `json:"expiration_date"`
}

type SomeGroceriesResponse struct {
	Groceries []*ent.Grocery `json:"groceries"`
}

type OneGroceriesResponse struct {
	Grocery *ent.Grocery `json:"grocery"`
}

func (c *GroceryController) GetAll() (*SomeGroceriesResponse, error) {
	groceries, err := c.GroceryService.GetAll(context.TODO())
	if err != nil {
		return nil, err
	}
	return &SomeGroceriesResponse{Groceries: groceries}, nil
}

func (c *GroceryController) Find(groceryId int) (*OneGroceriesResponse, error) {
	grocery, err := c.GroceryService.Find(context.TODO(), groceryId)
	if err != nil {
		return nil, err
	}
	return &OneGroceriesResponse{Grocery: grocery}, nil
}

func (c *GroceryController) Provide(req GroceryProvideRequest) (*OneGroceriesResponse, error) {
	grocery, err := c.GroceryService.Provide(context.TODO(), service.GroceryProvideRequest{
		ProviderId:     req.ProviderId,
		Name:           req.Name,
		Price:          req.Price,
		Unit:           req.Unit,
		ExpirationDate: req.ExpirationDate,
	})
	if err != nil {
		return nil, err
	}
	return &OneGroceriesResponse{Grocery: grocery}, nil
}
