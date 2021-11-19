package series

import "gorm.io/datatypes"

type SeriesResponse struct {
	ID         int            `json:"id"`
	Title      string         `json:"title"`
	Overview   string         `json:"overview"`
	PosterPath string         `json:"poster_path"`
	Popularity float64        `json:"popularity"`
	Tags       datatypes.JSON `json:"tags"`
}
