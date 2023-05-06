package routes

import (
	"bale-management/handlers"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	ticketGroup := router.Group("/tickets")
	{
		ticketGroup.GET("", handlers.GetTicketsHandler)
		ticketGroup.GET("/:id", handlers.GetTicketByIDHandler)
		ticketGroup.POST("", handlers.CreateTicketHandler)
		ticketGroup.PUT("/:id", handlers.UpdateTicketHandler)
		ticketGroup.DELETE("/:id", handlers.DeleteTicketHandler)
		ticketGroup.POST("/:id/mass", handlers.UpdateTicketMassHandler)
		ticketGroup.POST("/:id/price", handlers.UpdateTicketPriceHandler)
		ticketGroup.POST("/:id/barcode", handlers.UpdateTicketBarcodeHandler)
		ticketGroup.POST("/:id/green_grade", handlers.UpdateTicketGreenGradeHandler)
	}
}
