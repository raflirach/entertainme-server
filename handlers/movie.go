package handlers

import (
	"entertaime-server/models/movie"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type movieHandler struct {
	movieService movie.Service
}

func NewMovieHandle(movieService movie.Service) *movieHandler {
	return &movieHandler{movieService}
}

func (h *movieHandler) GetMovies(c *gin.Context) {
	movies, err := h.movieService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var moviesResponse []movie.MovieResponse

	for _, v := range movies {
		movieResponse := converToResponse(v)
		moviesResponse = append(moviesResponse, movieResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": moviesResponse,
	})

}

func (h *movieHandler) GetMoviesById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.movieService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	movieResponse := converToResponse(movie)

	c.JSON(http.StatusOK, gin.H{
		"movie": movieResponse,
	})

}

func (h *movieHandler) CreateMovie(c *gin.Context) {
	var request movie.MovieRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	movie, err := h.movieService.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})

}

func (h *movieHandler) UpdateMovie(c *gin.Context) {
	var request movie.MovieRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.movieService.Update(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})

}

func (h *movieHandler) DeleteMovie(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.movieService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Succesfully delete movie with id: %d", id),
		"movie":   movie,
	})

}

func converToResponse(m movie.Movie) movie.MovieResponse {
	return movie.MovieResponse{
		ID:         m.ID,
		Title:      m.Title,
		Overview:   m.Overview,
		PosterPath: m.PosterPath,
		Popularity: m.Popularity,
		Tags:       m.Tags,
	}
}
