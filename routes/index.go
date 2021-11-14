package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunRouter() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"github": "https://github.com/raflirach/entertainme-server",
		})
	})

	GetMovieRouter(router)

	router.Run()
}
