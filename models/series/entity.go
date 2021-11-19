package series

import (
	"time"

	"gorm.io/datatypes"
)

type Series struct {
	ID         int            `json:"id"`
	Title      string         `json:"title"`
	Overview   string         `json:"overview"`
	PosterPath string         `json:"poster_path"`
	Popularity float64        `json:"popularity"`
	Tags       datatypes.JSON `json:"tags"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
