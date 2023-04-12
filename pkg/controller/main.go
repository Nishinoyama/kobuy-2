package controller

import (
	"context"
	"errors"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/ent/ledger"
)

type SomeUsersResponse struct {
	Users []*ent.User `json:"users"`
}

func GetAllUsers(client *ent.UserClient, ctx context.Context) (*SomeUsersResponse, error) {
	users, err := client.Query().WithProvidedGroceries().WithPurchased().All(ctx)
	if err != nil {
		return nil, err
	}
	return &SomeUsersResponse{users}, nil
}

type OneUserResponse struct {
	User     *ent.User       `json:"user"`
	Purchase []*ent.Purchase `json:"purchase"`
}

func FindUser(client *ent.UserClient, ctx context.Context, userId int) (*OneUserResponse, error) {
	user, err := client.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	purchase, err := user.QueryPurchased().WithGrocery().All(ctx)
	if err != nil {
		return nil, err
	}
	return &OneUserResponse{user, purchase}, nil
}

func PurchaseGrocery(client *ent.Client, ctx context.Context, buyerId int, groceryId int, unit int) error {
	buyer, err := client.User.Get(ctx, buyerId)
	if err != nil {
		return err
	}
	grocery, err := client.Grocery.Get(ctx, groceryId)
	if err != nil {
		return err
	}
	seller, err := grocery.QueryProvider().First(ctx)
	if err != nil {
		return err
	}
	price := grocery.Price * unit

	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	if err := tx.Purchase.Create().SetPrice(grocery.Price).SetAmount(unit).SetBuyer(buyer).SetGrocery(grocery).Exec(ctx); err != nil {
		if tx.Rollback() != nil {
			return errors.New("roll back failed")
		}
		return err
	}
	if err := tx.Grocery.UpdateOne(grocery).SetUnit(grocery.Unit - unit).Exec(ctx); err != nil {
		if tx.Rollback() != nil {
			return errors.New("roll back failed")
		}
		return err
	}
	if err := tx.User.UpdateOne(buyer).AddBalance(-price).Exec(ctx); err != nil {
		if tx.Rollback() != nil {
			return errors.New("roll back failed")
		}
		return err
	}
	if err := tx.User.UpdateOne(seller).AddBalance(price).Exec(ctx); err != nil {
		if tx.Rollback() != nil {
			return errors.New("roll back failed")
		}
		return err
	}
	if err := tx.Ledger.Create().SetPayer(buyer).SetReceiver(seller).SetPrice(price).SetType(ledger.TypePurchase).Exec(ctx); err != nil {
		if tx.Rollback() != nil {
			return errors.New("roll back failed")
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
