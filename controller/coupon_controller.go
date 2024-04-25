package controller

import "github.com/gin-gonic/gin"

// GetAllCoupon godoc
// @Summary Get All Coupon
// @Description get All Coupon
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.CouponResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/coupon [get]
func GetAllCoupon(c *gin.Context) {

}
