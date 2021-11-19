package routes

import (
	"entertaime-server/config"
	"entertaime-server/handlers"
	"entertaime-server/models/series"

	"github.com/gin-gonic/gin"
)

func GetSeriesRouter(router *gin.Engine) {
	db := config.GetDatabase()

	seriesReporsitory := series.NewRepository(db)
	seriesService := series.NewService(seriesReporsitory)
	seriesHandler := handlers.NewSeriesHandle(seriesService)

	seriesRouter := router.Group("/series")

	seriesRouter.GET("/", seriesHandler.GetSeries)
	seriesRouter.GET("/:id", seriesHandler.GetSeriesById)
	seriesRouter.POST("/", seriesHandler.CreateSeries)
	seriesRouter.PUT("/:id", seriesHandler.UpdateSeries)
	seriesRouter.DELETE("/:id", seriesHandler.DeleteSeries)
}
