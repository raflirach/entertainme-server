package movie

type MovieRequest struct {
	Title      string  `json:"title" binding:"required"`
	Overview   string  `json:"overview"`
	PosterPath string  `json:"poster_path"`
	Popularity float64 `json:"popularity" binding:"required,number"`
	Tags       string  `json:"tags"`
}
