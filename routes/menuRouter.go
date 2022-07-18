package routes

import "github.com/gin-gonic/gin"

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenues())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenue())
	incomingRoutes.POST("/menus", controller.CreateMenue())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}
