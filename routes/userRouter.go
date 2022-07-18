package routes

import (
	controller "github/sajdakabir/restaurant_management_backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUP())
	incomingRoutes.POST("/users/login", controller.LOgin())
}
