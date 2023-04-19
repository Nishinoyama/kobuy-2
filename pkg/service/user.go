package service

import (
	"context"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
)

type UserService struct {
	UserClient *ent.UserClient
}

func (s *UserService) GetAll(ctx context.Context) ([]*ent.User, error) {
	return s.UserClient.Query().All(ctx)
}

func (s *UserService) Find(ctx context.Context, userId int) (*ent.User, error) {
	return s.UserClient.Get(ctx, userId)
}
