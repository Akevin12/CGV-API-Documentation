package controller

import (
	"github.com/gin-gonic/gin"
)

// GetPayments godoc
// @Summary Get Payments
// @Description Get all payments
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.TicketsResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/tickets [get]
func GetTickets(c *gin.Context) {
}

// CreateTicket godoc
// @Summary Create Ticket
// @Description Create Ticket by shedule
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.ScheduleTicket true "Schedule Detail"
// @Success 200 {object} models.TicketResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/tickets [post]
func CreateTicket(c *gin.Context) {
}

// GetPayment godoc
// @Summary Get Payment
// @Description Get payment
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param ticketId path int true "Customer ID"
// @Success 200 {object} models.TicketResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/tickets/{ticketId} [get]
func GetTicket(c *gin.Context) {
}

// ConfirmPayment godoc
// @Summary Confirm Payment
// @Description Confirm payment with payment_id.
// @Tags Customer
// @Param customerId path int true "Customer ID"
// @Param ticketId path string true "payment id"
// @Accept json
// @Produce json
// @Success 200 {object} models.TicketResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/tickets/{ticketId}/payment [post]
func ConfirmPayment(c *gin.Context) {
}
