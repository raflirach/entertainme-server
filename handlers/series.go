package handlers

import (
	"entertaime-server/models/series"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type seriesHandler struct {
	seriesService series.Service
}

func NewSeriesHandle(seriesService series.Service) *seriesHandler {
	return &seriesHandler{seriesService}
}

func (h *seriesHandler) GetSeries(c *gin.Context) {
	s, err := h.seriesService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var seriesResponse []series.SeriesResponse

	for _, v := range s {
		seriesResponse = append(seriesResponse, converToSeriesResponse(v))
	}

	c.JSON(http.StatusOK, gin.H{
		"series": seriesResponse,
	})

}

func (h *seriesHandler) GetSeriesById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	series, err := h.seriesService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	seriesResponse := converToSeriesResponse(series)

	c.JSON(http.StatusOK, gin.H{
		"series": seriesResponse,
	})

}

func (h *seriesHandler) CreateSeries(c *gin.Context) {
	var request series.SeriesRequest

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

	series, err := h.seriesService.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"series": series,
	})

}

func (h *seriesHandler) UpdateSeries(c *gin.Context) {
	var request series.SeriesRequest

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

	series, err := h.seriesService.Update(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"series": series,
	})

}

func (h *seriesHandler) DeleteSeries(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	series, err := h.seriesService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Succesfully delete series with id: %d", id),
		"series":  series,
	})

}

func converToSeriesResponse(s series.Series) series.SeriesResponse {
	return series.SeriesResponse{
		ID:         s.ID,
		Title:      s.Title,
		Overview:   s.Overview,
		PosterPath: s.PosterPath,
		Popularity: s.Popularity,
		Tags:       s.Tags,
	}
}
