package movie

import "time"

type Movie struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Overview   string  `json:"overview"`
	PosterPath string  `json:"poster_path"`
	Popularity float64 `json:"popularity"`
	Tags       string  `json:"tags"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
