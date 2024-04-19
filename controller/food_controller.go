package controller

import (
	"github.com/gin-gonic/gin"
)

// Create Add Chart godoc
// @Summary Add To Chart
// @Description Add To Chart
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.Chart true "Chart details"
// @Success 200 {object} models.ChartResponse
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
// @Router /customer/{customerId}/chart [delete]
func DeleteChart(c *gin.Context) {
}
