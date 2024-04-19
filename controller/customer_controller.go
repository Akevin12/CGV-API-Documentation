package controller

import (
	"github.com/gin-gonic/gin"
)

// CreateCustomer godoc
// @Summary Create New Customer
// @Description Create New Customer Account
// @Tags Customer
// @Accept json
// @Produce json
// @Param body body models.Customer true "Customer details"
// @Success 200 {object} models.CustomerResponse
// @Router /customer/registration [post]
func AddCustomer(c *gin.Context) {
}

// LoginCustomer godoc
// @Summary Login Customer
// @Description Login Customer Account
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Login details"
// @Success 200 {object} models.CustomerResponse
// @Router /customer/auth/login [post]
func LoginCustomer(c *gin.Context) {
}

// LogoutAccount godoc
// @Summary Logout Account
// @Description Logout Account admin and customer
// @Tags Auth
// @Success 200 {object} models.Response
// @Router /auth/logout [delete]
func LogoutAccount(c *gin.Context) {
}

// GetCustomer godoc
// @Summary Get Customer
// @Description get Account
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.CustomerResponse
// @Router /customer/{customerId}/profile [get]
func GetCustomer(c *gin.Context) {
}

// PutCustomer godoc
// @Summary Update Customer
// @Description Put Customer Account
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.Customer true "Customer details"
// @Success 200 {object} models.CustomerResponse
// @Router /customer/{customerId}/profile [put]
func UpdateCustomer(c *gin.Context) {
}

// CreategetMemberById godoc
// @Summary Get Member By Id
// @Description Get Member By Id
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.MemberResponse
// @Router /customer/{customerId}/member [get]
func getMemberById(c *gin.Context) {
}
