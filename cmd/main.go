package main

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"github.com/nishinoyama/kobuy-2/pkg/http/handler"
	"github.com/nishinoyama/kobuy-2/pkg/service"
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

	userService := service.UserService{UserClient: client.User}
	groceryService := service.GroceryService{GroceryClient: client.Grocery}
	ledgerService := service.LedgerService{LedgerClient: client.Ledger}

	userController := controller.UserController{UserService: &userService}
	groceryController := controller.GroceryController{GroceryService: &groceryService}
	ledgerController := controller.LedgerController{LedgerService: ledgerService}

	engine := gin.Default()

	v1 := engine.Group("/v1/api")
	{
		handler.NewUserHandler(v1, &userController)
		handler.NewGroceryHandler(v1, &groceryController)
		handler.NewLedgerHandler(v1, &ledgerController)

		handler.NewPurchaseGroceryHandler(v1, client)
		handler.NewCashHandler(v1, client)

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
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
