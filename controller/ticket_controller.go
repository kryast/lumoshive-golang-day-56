package controller

import (
	"day-56/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TicketController struct {
	repo   repository.TicketRepository
	logger *zap.Logger
}

func NewTicketController(db *gorm.DB) *TicketController {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize zap logger: " + err.Error())
	}

	return &TicketController{
		repo:   repository.NewTicketRepository(db),
		logger: logger,
	}
}

func (ctrl *TicketController) All(c *gin.Context) {
	tickets, err := ctrl.repo.All()
	if err != nil {
		ctrl.logger.Error("Failed to retrieve tickets", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	ctrl.logger.Info("Successfully retrieved tickets", zap.Int("count", len(tickets)))
	c.JSON(http.StatusOK, gin.H{"products": tickets})
}

func (ctrl *TicketController) Update(c *gin.Context) {
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Warn("Invalid ticket ID", zap.String("id", c.Param("id")), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	var request struct {
		Quantity uint `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		ctrl.logger.Warn("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctrl.logger.Info("Attempting to update ticket quantity", zap.Int("ticketID", ticketID), zap.Uint("newQuantity", request.Quantity))

	if err := ctrl.repo.Update(uint(ticketID), request.Quantity); err != nil {
		ctrl.logger.Error("Failed to update ticket quantity", zap.Int("ticketID", ticketID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quantity"})
		return
	}

	ctrl.logger.Info("Successfully updated ticket quantity", zap.Int("ticketID", ticketID), zap.Uint("newQuantity", request.Quantity))
	c.JSON(http.StatusOK, gin.H{"message": "Quantity updated successfully"})
}
