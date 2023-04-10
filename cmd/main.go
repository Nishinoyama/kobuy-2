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
	"time"
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
		v1.GET("/users", getUsersHandler(client.User))
		v1.GET("/groceries", getGroceriesHandler(client.Grocery))
		v1.POST("/groceries/provide", provideGroceryHandler(client.Grocery))
	}
	if err := engine.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}

func getUsersHandler(userClient *ent.UserClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		users, err := userClient.Query().All(cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, users)
	}
}

func getGroceriesHandler(groceryClient *ent.GroceryClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		users, err := groceryClient.Query().WithProvider().All(cc)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		gc.JSON(http.StatusOK, users)
	}
}

type ProvideGroceryRequest struct {
	ProviderId     int        `json:"provider_id,omitempty" binding:"required"`
	Name           string     `json:"name,omitempty" binding:"required"`
	Price          int        `json:"price,omitempty" binding:"required"`
	Unit           int        `json:"unit,omitempty" binding:"required"`
	ExpirationDate *time.Time `json:"expiration_date"`
}

func provideGroceryHandler(groceryClient *ent.GroceryClient) func(ctx *gin.Context) {
	return func(gc *gin.Context) {
		cc := context.Background()
		var req ProvideGroceryRequest
		if err := gc.BindJSON(&req); err != nil {
			gc.JSON(http.StatusBadRequest, err.Error())
			return
		}
		g := groceryClient.Create().
			SetProviderID(req.ProviderId).
			SetName(req.Name).
			SetPrice(req.Price).
			SetUnit(req.Unit)
		if req.ExpirationDate != nil {
			g.SetExpirationDate(*req.ExpirationDate)
		}
		if grocery, err := g.Save(cc); err != nil {
			gc.JSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			gc.JSON(http.StatusOK, grocery)
			return
		}
	}
}
