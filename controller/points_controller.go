package controller

import "github.com/gin-gonic/gin"

// GetPointsHistory godoc
// @Summary Get Points History
// @Description Get Points History
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.PointsResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/points [get]
func GetPointsHistory(c *gin.Context) {

}
