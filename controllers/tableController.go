package controller

import (
	"context"
	"github/sajdakabir/restaurant_management_backend/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := tableCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when listing tables"})
			return
		}
		var allTables []bson.M
		if err = result.All(ctx, allTables); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allTables)

	}
}
