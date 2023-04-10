package main

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/pkg/handler"
	"log"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
	faker.SetGenerateUniqueValues(true)
	if err := client.User.CreateBulk(
		client.User.Create().SetName("taro"),
		client.User.Create().SetName("jiro"),
		client.User.Create().SetName("saro"),
	).Exec(context.TODO()); err != nil {
		log.Fatal(err)
	}
	engine := gin.Default()
	v1 := engine.Group("/v1/api")
	{
		v1.GET("/users", handler.GetUsersHandler(client.User))
		v1.GET("/groceries", handler.GetGroceriesHandler(client.Grocery))
		v1.POST("/groceries/provide", handler.ProvideGroceryHandler(client.Grocery))
	}
	if err := engine.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
