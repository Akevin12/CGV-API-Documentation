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
// @Success 200 {object} models.MoviesResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /movies/search [get]
func SearchMovies(c *gin.Context) {
}

// SearchMovies godoc
// @Summary Search cinema
// @Description Search cinema
// @Tags Guest
// @Accept json
// @Produce json
// @Param name query string false "Cinema name to search"
// @Success 200 {object} models.BranchSearchResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /branch/search [get]
func SearchCinema(c *gin.Context) {
}

// Add To Favorite godoc
// @Summary Add To Favorite
// @Description Add To Favorite
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.BranchDetail true "Branch details"
// @Success 200 {object} models.BranchDetailResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/favorite [post]
func AddToFavorite(c *gin.Context) {
}

// Delete Favorite godoc
// @Summary Delete Favorite
// @Description Delete Favorite
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/favorite [delete]
func DeleteFavorite(c *gin.Context) {
}

// AddWishlist godoc
// @Summary Add Watchlist
// @Description Add Watchlist
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.MovieSchedule true "Watchlist details"
// @Success 200 {object} models.MovieSchedulesResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/watchlist [post]
func AddWatchlist(c *gin.Context) {
}

// getWishlist godoc
// @Summary Get Watchlist
// @Description Get Watchlist
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} models.MovieSchedulesResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/watchlist [get]
func GetWatchlist(c *gin.Context) {
}

// DeleteWatchlist godoc
// @Summary Delete Watchlist
// @Description Delete Watchlist
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param WatchlistId path int true "Watchlist ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/watchlist/{watchlistId} [delete]
func DeleteWatchlist(c *gin.Context) {
}

// GetAllReviews godoc
// @Summary Get All Reviews
// @Description Get All Reviews
// @Tags Guest
// @Accept json
// @Produce json
// @Param movieId path int true "Movie ID"
// @Success 200 {object} models.ReviewsResponse
// @Failure 400 {object} models.ErrorResponse "Error : Couldn't get reviews!"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /movie/reviews/{movieId} [get]
func GetAllReviews(c *gin.Context) {
}

// AddReview godoc
// @Summary Add Review
// @Description Add Review
// @Tags Customer
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID"
// @Param body body models.ReviewFormBody true "Review details"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse "Error : Couldn't add review!"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /customer/{customerId}/movie/review [post]
func AddReview(c *gin.Context) {
}

// Now Showing Movies godoc
// @Summary Get Now Showing Movies
// @Description Get Now Showing Movies
// @Tags Guest
// @Accept json
// @Produce json
// @Param body body models.NowShowingMovies true "Now Showing Movies"
// @Success 200 {object} models.NowShowingMoviesResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /movies/nowshowing [get]
func NowShowing(c *gin.Context) {
}

// Upcoming Movies godoc
// @Summary Get Upcoming Movies
// @Description Get Upcoming Movies
// @Tags Guest
// @Accept json
// @Produce json
// @Param body body models.UpComingMovies true "Upcoming Movies"
// @Success 200 {object} models.UpComingMoviesResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /movies/upcoming [get]
func Upcoming(c *gin.Context) {
}
