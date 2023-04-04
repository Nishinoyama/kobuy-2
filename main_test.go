package main

import (
	"context"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/stretchr/testify/assert"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func TestMainKobuy2InMemory(t *testing.T) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	faker.SetGenerateUniqueValues(true)
	{
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
				groceryCreates = append(groceryCreates, client.Grocery.Create().SetName(faker.Word()).SetProvider(user))
			}
		}
		if err := client.Grocery.CreateBulk(groceryCreates...).Exec(ctx); err != nil {
			t.Fatal(err)
		}
		cnt := client.Grocery.Query().CountX(ctx)
		if !assert.Equal(t, cnt, 15, "3 * 5 !=", cnt) {
			t.FailNow()
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
			t.Log(g.Name, "by", u.Name)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	// serialize
	marshal, err := json.Marshal(us)
	t.Logf("%s", marshal)
}
