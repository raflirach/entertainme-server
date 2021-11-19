package series

import "gorm.io/datatypes"

type SeriesRequest struct {
	Title      string         `json:"title" binding:"required"`
	Overview   string         `json:"overview"`
	PosterPath string         `json:"poster_path"`
	Popularity float64        `json:"popularity" binding:"required,number"`
	Tags       datatypes.JSON `json:"tags"`
}
