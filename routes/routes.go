package routes

import (
	"CGV/controller"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/information/news", controller.GetAllNews)
	router.GET("/information/promotions", controller.GetAllPromotions)
	router.GET("/information/features", controller.GetAllFeatures)
	router.GET("/movie/reviews/:movieId", controller.GetAllReviews)
	return router
}
