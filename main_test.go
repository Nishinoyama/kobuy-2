package main

import (
	"context"
	"encoding/json"
	"entgo.io/ent/dialect"
	"github.com/go-faker/faker/v4"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"github.com/nishinoyama/kobuy-2/pkg/ent/ledger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainKobuy2InMemory(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	//client, err := ent.Open(dialect.MySQL, "root:pass@tcp(localhost:3306)/test?parseTime=True")
	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	{
		faker.SetGenerateUniqueValues(true)
		userCreates := make([]*ent.UserCreate, 5)
		for i := 0; i < 5; i++ {
			userCreates[i] = client.User.Create().SetName(faker.FirstName())
		}
		users, err := client.User.CreateBulk(userCreates...).Save(ctx)
		if err != nil {
			t.Fatal(err)
		}

		groceryCreates := make([]*ent.GroceryCreate, 0, 15)
		for _, user := range users {
			for i := 0; i < 3; i++ {
				groceryCreates = append(
					groceryCreates,
					client.Grocery.Create().SetName(faker.Word()).SetPrice(100).SetUnit(i*3+1).SetProvider(user),
				)
			}
		}
		_, err = client.Grocery.CreateBulk(groceryCreates...).Save(ctx)
		if err != nil {
			t.Fatal(err)
		}

		{
			buyer := client.User.GetX(ctx, 2)
			g := client.Grocery.GetX(ctx, 1)
			seller := client.Grocery.QueryProvider(g).OnlyX(ctx)

			tx, err := client.Tx(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if err := tx.Purchase.Create().SetPrice(g.Price).SetAmount(1).SetBuyer(buyer).SetGrocery(g).Exec(ctx); err != nil {
				if tx.Rollback() != nil {
					t.Fatal("what")
				}
				t.Fatal(err)
			}
			if err := tx.Grocery.UpdateOne(g).SetUnit(g.Unit - 1).Exec(ctx); err != nil {
				if tx.Rollback() != nil {
					t.Fatal("what")
				}
				t.Fatal(err)
			}
			if err := tx.User.UpdateOne(buyer).SetBalance(buyer.Balance - g.Price).Exec(ctx); err != nil {
				if tx.Rollback() != nil {
					t.Fatal("what")
				}
				t.Fatal(err)
			}
			if err := tx.User.UpdateOne(seller).SetBalance(seller.Balance + g.Price).Exec(ctx); err != nil {
				if tx.Rollback() != nil {
					t.Fatal("what")
				}
				t.Fatal(err)
			}
			if err := tx.Ledger.Create().SetPayer(buyer).SetReceiver(seller).SetPrice(g.Price).SetType(ledger.TypePurchase).Exec(ctx); err != nil {
				if tx.Rollback() != nil {
					t.Fatal("what")
				}
				t.Fatal(err)
			}
			if err := tx.Commit(); err != nil {
				t.Fatal(err)
			}
		}
		{
			buyer := client.User.GetX(ctx, 2)
			g := client.Grocery.GetX(ctx, 1)
			seller := client.Grocery.QueryProvider(g).OnlyX(ctx)
			assert.Less(t, buyer.Balance, 0)
			assert.Greater(t, seller.Balance, 0)
		}
	}

	//// N + 1
	//for _, g := range client.Debug().Grocery.Query().AllX(ctx) {
	//	provider := g.QueryProvider().FirstX(ctx)
	//	t.Log(g.Name, "by", provider.Name)
	//}

	for _, g := range client.Debug().Grocery.Query().WithProvider().AllX(ctx) {
		t.Log(g.Name, "by", g.Edges.Provider.Name)
	}

	us := client.Debug().User.Query().WithProvidedGroceries().AllX(ctx)
	for _, u := range us {
		for _, g := range u.Edges.ProvidedGroceries {
			t.Log(g.Name, "by", u.Name, "rest", g.Unit, "price", g.Price, "expire at", g.ExpirationDate)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	// serialize
	marshal, err := json.Marshal(us)
	t.Logf("%s", marshal)

}
