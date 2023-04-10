package controller

import (
	"context"
	"github.com/nishinoyama/kobuy-2/ent"
)

type SomeUsersResponse struct {
	Users []*ent.User `json:"users"`
}

func GetAllUsers(client *ent.UserClient, ctx context.Context) (*SomeUsersResponse, error) {
	users, err := client.Query().WithProvidedGroceries().All(ctx)
	if err != nil {
		return nil, err
	}
	return &SomeUsersResponse{users}, nil
}

type OneUserResponse struct {
	User *ent.User `json:"user"`
}

func FindUser(client *ent.UserClient, ctx context.Context, userId int) (*OneUserResponse, error) {
	user, err := client.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &OneUserResponse{user}, nil
}
