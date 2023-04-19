package service

import (
	"context"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"time"
)

type GroceryService struct {
	GroceryClient *ent.GroceryClient
}

type GroceryProvideRequest struct {
	ProviderId     int
	Name           string
	Price          int
	Unit           int
	ExpirationDate *time.Time
}

func (s *GroceryService) GetAll(ctx context.Context) ([]*ent.Grocery, error) {
	return s.GroceryClient.Query().WithProvider().All(ctx)
}

func (s *GroceryService) Find(ctx context.Context, groceryId int) (*ent.Grocery, error) {
	return s.GroceryClient.Get(ctx, groceryId)
}

func (s *GroceryService) Provide(ctx context.Context, req GroceryProvideRequest) (*ent.Grocery, error) {
	g := s.GroceryClient.Create().
		SetProviderID(req.ProviderId).
		SetName(req.Name).
		SetPrice(req.Price).
		SetUnit(req.Unit).
		SetNillableExpirationDate(req.ExpirationDate)
	return g.Save(ctx)
}
