package repository

import (
	"day-56/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	All() ([]models.Ticket, error)
	Update(ticketID uint, newQuantity uint) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) All() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) Update(ticketID uint, newQuantity uint) error {
	result := r.db.Model(&models.Ticket{}).Where("id = ?", ticketID).Update("quantity", newQuantity)
	return result.Error
}
