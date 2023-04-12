package controller

import (
	"context"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/pkg/service"
)

type UserController struct {
	UserService *service.UserService
}

type SomeUsersResponse struct {
	Users []*ent.User `json:"users"`
}

type OneUserResponse struct {
	User *ent.User `json:"user"`
}

func (c *UserController) GetAll() (*SomeUsersResponse, error) {
	users, err := c.UserService.GetAll(context.TODO())
	if err != nil {
		return nil, err
	}
	return &SomeUsersResponse{users}, nil
}

func (c *UserController) Find(userId int) (*OneUserResponse, error) {
	user, err := c.UserService.Find(context.TODO(), userId)
	if err != nil {
		return nil, err
	}
	return &OneUserResponse{user}, nil
}
