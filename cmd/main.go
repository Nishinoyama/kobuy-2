package main

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"github.com/nishinoyama/kobuy-2/pkg/http/handler"
	"github.com/nishinoyama/kobuy-2/pkg/service"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

type Configure struct {
	CorsOrigins []string `yaml:"corsOrigins"`
}

func main() {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var config Configure
	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatal(err)
	}
	if err := configFile.Close(); err != nil {
		log.Fatal(err)
	}

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
			client.User.Create().SetName("taro").SetPassword("taro"),
			client.User.Create().SetName("jiro").SetPassword("jiro"),
			client.User.Create().SetName("saro").SetPassword("saro"),
		).Exec(context.TODO()); err != nil {
			log.Fatal(err)
		}
		if err := client.Grocery.CreateBulk(
			client.Grocery.Create().SetProviderID(1).SetName("choco").SetPrice(120).SetUnit(23).SetExpirationDate(time.Now().AddDate(10, 0, 0)),
			client.Grocery.Create().SetProviderID(1).SetName("vanilla").SetPrice(110).SetUnit(23).SetExpirationDate(time.Now().AddDate(0, -1, 0)),
			client.Grocery.Create().SetProviderID(2).SetName("mow").SetPrice(170).SetUnit(8).SetExpirationDate(time.Now().AddDate(0, 0, 3)),
			client.Grocery.Create().SetProviderID(2).SetName("shiro").SetPrice(320).SetUnit(3),
			client.Grocery.Create().SetProviderID(1).SetName("smile").SetPrice(0).SetUnit(9999).SetExpirationDate(time.Now().AddDate(1000, 0, 0)),
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
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     config.CorsOrigins,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		MaxAge: 24 * time.Hour,
	}))

	handler.LoginHandler(engine, client.User)

	v1 := engine.Group("/v1/api")
	{
		handler.NewUserHandler(v1, &userController)
		handler.NewGroceryHandler(v1, &groceryController)
		handler.NewLedgerHandler(v1, &ledgerController)

		handler.NewPurchaseHandler(v1, client)
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
