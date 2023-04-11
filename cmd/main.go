package main

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/pkg/handler"
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

	{
		// seeding
		if err := client.User.CreateBulk(
			client.User.Create().SetName("taro"),
			client.User.Create().SetName("jiro"),
			client.User.Create().SetName("saro"),
		).Exec(context.TODO()); err != nil {
			log.Fatal(err)
		}
		if err := client.Grocery.CreateBulk(
			client.Grocery.Create().SetProviderID(1).SetName("choco").SetPrice(120).SetUnit(23),
			client.Grocery.Create().SetProviderID(1).SetName("vanilla").SetPrice(110).SetUnit(23),
			client.Grocery.Create().SetProviderID(2).SetName("mow").SetPrice(170).SetUnit(8),
			client.Grocery.Create().SetProviderID(2).SetName("shiro").SetPrice(320).SetUnit(3),
		).Exec(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}

	engine := gin.Default()
	v1 := engine.Group("/v1/api")
	{
		v1.GET("/users", handler.GetUsersHandler(client.User))
		v1.GET("/users/:id", handler.FindUserHandler(client.User))
		v1.GET("/groceries", handler.GetGroceriesHandler(client.Grocery))
		v1.GET("/ledger", handler.GetLedger(client))
		v1.POST("/groceries/provide", handler.ProvideGroceryHandler(client.Grocery))
		v1.POST("/groceries/purchase", handler.PurchaseGroceryHandler(client))
		v1.POST("/ledger/cash", handler.CashLedgerHandler(client))
	}
	{
		engine.LoadHTMLGlob("./cmd/templates/*")
		engine.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "taro",
			})
		})
		engine.GET("/groceries", func(c *gin.Context) {
			groceries, err := client.Grocery.Query().WithProvider().All(context.TODO())
			if err != nil {
				c.HTML(http.StatusInternalServerError, "500.tmpl",
					struct{ mes string }{err.Error()},
				)
				return
			}
			c.HTML(http.StatusOK, "groceries.tmpl", gin.H{
				"groceries": groceries,
			})
		})
	}
	if err := engine.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
