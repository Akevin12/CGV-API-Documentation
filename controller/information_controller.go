package controller

import (
	"github.com/gin-gonic/gin"
)

// GetAllNews godoc
// @Summary Get all news
// @Description Get all news
// @Tags Guest
// @Accept json
// @Produce json
// @Success 200 {object} models.NewsArrayResponse
// @Failure 400 {object} models.ErrorResponse "Error : Couldn't get all news!"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /information/news [get]
func GetAllNews(c *gin.Context) {
}

// GetAllPromotions godoc
// @Summary Get All Promotions
// @Description Get All Promotions
// @Tags Guest
// @Accept json
// @Produce json
// @Success 200 {object} models.PromotionsResponse
// @Failure 400 {object} models.ErrorResponse "Error : Couldn't get promotions!"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /information/promotions [get]
func GetAllPromotions(c *gin.Context) {
}

// GetAllFeatures godoc
// @Summary Get All Features
// @Description Get All Features
// @Tags Guest
// @Accept json
// @Produce json
// @Success 200 {object} models.FeaturesResponse
// @Failure 400 {object} models.ErrorResponse "Error : Couldn't get Features!"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /information/features [get]
func GetAllFeatures(c *gin.Context) {
}
