package routes

import (
	"day-56/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TicketRoutes(router *gin.Engine, db *gorm.DB) {
	ticketController := controller.NewTicketController(db)

	ticketRoutes := router.Group("/ticket")
	{
		ticketRoutes.GET("/", ticketController.All)
		ticketRoutes.PUT("/:id", ticketController.Update)

	}
}
