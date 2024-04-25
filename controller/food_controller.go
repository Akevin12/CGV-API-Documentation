package controller

import (
	"github.com/gin-gonic/gin"
)

// Create Add Chart godoc
// @Summary Get Chart
// @Description Get Chart
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.ChartResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/chart [get]
func getChart(c *gin.Context) {
}

// Create Add Chart godoc
// @Summary Add To Chart
// @Description Add To Chart
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.Chart true "Chart details"
// @Success 200 {object} models.ChartResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/chart [post]
func AddToChart(c *gin.Context) {
}

// Create Update Chart godoc
// @Summary Update Chart
// @Description Update Chart
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.Chart true "Chart details"
// @Success 200 {object} models.ChartResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/chart [put]
func UpdateChart(c *gin.Context) {
}

// Create Delete Chart godoc
// @Summary Delete Chart
// @Description Delete Chart
// @Tags Customer
// @Accept json
// @Produce json
// @Param ChartId path int true "Chart ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/chart [delete]
func DeleteChart(c *gin.Context) {
}

// Create Create Order Food godoc
// @Summary Create Order Food
// @Description Create Order Food
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.OrderFoodResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/orderfood [post]
func CreateOrderFood(c *gin.Context) {
}

// Create Get Food Voucher godoc
// @Summary Get Food Voucher
// @Description Get Food Voucher
// @Tags Customer
// @Accept json
// @Produce json
// @Param voucherName path string true "Name"
// @Param customerID path int true "Customer ID"
// @Success 200 {object} models.VoucherResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/foodvoucher [get]
func GetFoodVoucher(c *gin.Context) {
}

// Cancel Order Food godoc
// @Summary Cancel Order Food
// @Description Cancel Order Food
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerID path int true "Customer ID"
// @Param orderID path int true "Order ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/cancelorderfood [delete]
func CancelOrderFood(c *gin.Context) {
}
