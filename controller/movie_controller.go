package controller

import (
	"github.com/gin-gonic/gin"
)

// SearchMovies godoc
// @Summary Search movies by title
// @Description Search movies by title
// @Tags Guest
// @Accept json
// @Produce json
// @Param title query string false "Movie title to search"
// @Param genre query string false "Movie genre to search"
// @Success 200 {object} models.MoviesResponse
// @Router /movies/search [get]
func SearchMovies(c *gin.Context) {
}

// AddWishlist godoc
// @Summary Add Wishlist
// @Description Add Wishlist
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.MovieSchedule true "Whislist details"
// @Success 200 {object} models.MovieSchedulesResponse
// @Router /customer/{customerId}/wishlist [post]
func AddWishlist(c *gin.Context) {
}

// getWishlist godoc
// @Summary Get Wishlist
// @Description Get Wishlist
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.MovieSchedulesResponse
// @Router /customer/{customerId}/wishlist [get]
func GetWishlist(c *gin.Context) {
}

// DeleteWishlist godoc
// @Summary Delete Wishlist
// @Description Delete Wishlist
// @Tags Customer
// @Accept json
// @Produce json
// @Param WishlistId path int true "Wishlist ID"
// @Success 200 {object} models.Response
// @Router /customer/{customerId}/wishlist [delete]
func DeleteWishlist(c *gin.Context) {
}
