package movie

import "gorm.io/datatypes"

type MovieRequest struct {
	Title      string         `json:"title" binding:"required"`
	Overview   string         `json:"overview"`
	PosterPath string         `json:"poster_path"`
	Popularity float64        `json:"popularity" binding:"required,number"`
	Tags       datatypes.JSON `json:"tags"`
}
