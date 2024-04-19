package controller

import (
	"github.com/gin-gonic/gin"
)

// GetSchedule godoc
// @Summary Get Schedule
// @Description Get a schedule by movie_id and schedule_id.
// @Tags Customer
// @Param movieId path string true "movie id"
// @Accept json
// @Produce json
// @Success 200 {object} models.MovieSchedulesResponse
// @Router /movies/{movieId}/schedules [get]
func GetSchedules(c *gin.Context) {
}

// GetSchedulesByMovieId godoc
// @Summary Get all schedules for a movie by ID
// @Description Get all schedules for a movie by ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param movieId path string true "Movie ID"
// @Param scheduleId path string true "Schedule ID"
// @Success 200 {object} models.MovieScheduleResponse
// @Router /movies/{movieId}/schedules/{scheduleId} [get]
func GetSchedule(c *gin.Context) {
}
