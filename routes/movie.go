package routes

import (
	"entertaime-server/config"
	"entertaime-server/handlers"
	"entertaime-server/models/movie"

	"github.com/gin-gonic/gin"
)

func GetMovieRouter(router *gin.Engine) {
	db := config.GetDatabase()

	movieReporsitory := movie.NewRepository(db)
	movieService := movie.NewService(movieReporsitory)
	movieHandler := handlers.NewMovieHandle(movieService)

	movieRouter := router.Group("/movies")

	movieRouter.GET("/", movieHandler.GetMovies)
	movieRouter.GET("/:id", movieHandler.GetMoviesById)
	movieRouter.POST("/", movieHandler.CreateMovie)
	movieRouter.PUT("/:id", movieHandler.UpdateMovie)
	movieRouter.DELETE("/:id", movieHandler.DeleteMovie)
}
