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
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
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
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/auth/login [post]
func LoginCustomer(c *gin.Context) {
}

// LogoutAccount godoc
// @Summary Logout Account
// @Description Logout Account admin and customer
// @Tags Auth
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
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
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
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
// @Param body body models.CustomerEdit true "Customer details"
// @Success 200 {object} models.CustomerResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/profile [put]
func UpdateCustomer(c *gin.Context) {
}

// PutCustomer godoc
// @Summary Update Password
// @Description Put Password
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.Password true "Password details"
// @Success 200 {object} models.PassResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/password [put]
func UpdatePassword(c *gin.Context) {
}

// CreategetMemberById godoc
// @Summary Get Member
// @Description Get Member
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.MemberResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/member [get]
func getMemberById(c *gin.Context) {
}
