package controller

import (
	"net/http"

	m "CGV/models"

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
	// Placeholder implementation for fetching news
	placeholderNews := []m.News{
		{ID: 1, Title: "News Title 1", PicturePath: "/images/news1.jpg", Link: "https://example.com/news/1", Date: "2024-04-25"},
		{ID: 2, Title: "News Title 2", PicturePath: "/images/news2.jpg", Link: "https://example.com/news/2", Date: "2024-04-24"},
	}

	// Wrap the news array in a NewsArray struct
	newsArray := m.NewsArray{ArrNews: placeholderNews}

	// Wrap the NewsArray in a NewsArrayResponse along with a generic response
	response := m.NewsArrayResponse{
		Response: m.Response{
			Status:  http.StatusOK,
			Message: "Success",
		},
		NewsArray: newsArray,
	}

	c.JSON(http.StatusOK, response)
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
	// Placeholder implementation for fetching promotions
	placeholderPromotions := []m.Promotion{
		{ID: 1, Title: "Beli Combo Solo Tambah Mulai RP 5000* Dapat Merchandise?!", Category: "Concession", PicturePath: "/images/promotion1.jpg", Link: "https://example.com/promotion/1", Date: "2024-04-25"},
		{ID: 2, Title: "Nonton Hemat Pake QRIS Jago Cashback 40%!", Category: "Cinema", PicturePath: "/images/promotion2.jpg", Link: "https://example.com/promotion/2", Date: "2024-04-24"},
	}

	// Wrap the promotions array in a Promotions struct
	promotions := m.Promotions{Promotions: placeholderPromotions}

	// Wrap the Promotions struct in a PromotionsResponse along with a generic response
	response := m.PromotionsResponse{
		Response:   m.Response{Status: http.StatusOK, Message: "Success"},
		Promotions: promotions,
	}

	c.JSON(http.StatusOK, response)
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
	// Placeholder implementation for fetching features
	placeholderFeatures := []m.Feature{
		{ID: 1, Title: "Gold Class", Category: "Auditoriums", PicturePath: "/images/feature1.jpg", Link: "https://example.com/feature/1"},
		{ID: 2, Title: "SportsHall FX Sudirman", Category: "Sports", PicturePath: "/images/feature2.jpg", Link: "https://example.com/feature/2"},
	}

	// Wrap the features array in a Features struct
	features := m.Features{Features: placeholderFeatures}

	// Wrap the Features struct in a FeaturesResponse along with a generic response
	response := m.FeaturesResponse{
		Response: m.Response{Status: http.StatusOK, Message: "Success"},
		Features: features,
	}

	c.JSON(http.StatusOK, response)
}
