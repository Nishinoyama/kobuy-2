package service

import (
	"context"
	"errors"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"github.com/nishinoyama/kobuy-2/pkg/ent/ledger"
)

func Purchase(ctx context.Context, client *ent.Client, buyerId int, groceryId int, unit int) error {
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

func Cash(ctx context.Context, client *ent.Client, donorId int, receiverId int, price int) (*ent.Ledger, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.User.UpdateOneID(donorId).AddBalance(-price).Exec(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}
	if err := tx.User.UpdateOneID(receiverId).AddBalance(price).Exec(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}
	l, err := tx.Ledger.Create().
		SetPayerID(donorId).
		SetReceiverID(receiverId).
		SetPrice(price).
		SetType(ledger.TypeCash).Save(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return l, nil
}
