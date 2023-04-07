package main

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nishinoyama/kobuy-2/ent"
	"log"
	"net/http"
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
		client.User.Create().SetName(faker.FirstName()),
		client.User.Create().SetName(faker.FirstName()),
		client.User.Create().SetName(faker.FirstName()),
	).Exec(context.TODO()); err != nil {
		log.Fatal(err)
	}
	engine := gin.Default()
	engine.GET("/users", h(client))
	if err := engine.Run("localhost:3141"); err != nil {
		log.Fatal(err)
	}
}

func h(client *ent.Client) func(ctx *gin.Context) {
	cctx := context.Background()
	return func(ctx *gin.Context) {
		users, err := client.User.Query().All(cctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}
