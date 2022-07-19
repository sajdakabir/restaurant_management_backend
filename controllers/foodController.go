package controller

import (
	"context"
	"github/sajdakabir/restaurant_management_backend/database"
	"github/sajdakabir/restaurant_management_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while featching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}
