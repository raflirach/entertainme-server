package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func RunRouter() {

	router := gin.Default()
	godotenv.Load()
	os.Getenv("PORT")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"github": "https://github.com/raflirach/entertainme-server",
		})
	})

	GetMovieRouter(router)
	GetSeriesRouter(router)

	router.Run()
}
