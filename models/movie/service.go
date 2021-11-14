package movie

type Service interface {
	FindAll() ([]Movie, error)
	FindById(ID int) (Movie, error)
	Create(movieRequest MovieRequest) (Movie, error)
	Update(ID int, movieRequest MovieRequest) (Movie, error)
	Delete(ID int) (Movie, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Movie, error) {
	movies, err := s.repository.FindAll()
	return movies, err
}

func (s *service) FindById(ID int) (Movie, error) {
	movie, err := s.repository.FindById(ID)
	return movie, err
}

func (s *service) Create(movieRequest MovieRequest) (Movie, error) {
	movie := Movie{
		Title:      movieRequest.Title,
		Overview:   movieRequest.Overview,
		PosterPath: movieRequest.PosterPath,
		Popularity: movieRequest.Popularity,
		Tags:       movieRequest.Tags,
	}

	newMovie, err := s.repository.Create(movie)
	return newMovie, err
}

func (s *service) Update(ID int, movieRequest MovieRequest) (Movie, error) {
	movie, err := s.repository.FindById(ID)

	movie.Title = movieRequest.Title
	movie.Overview = movieRequest.Overview
	movie.PosterPath = movieRequest.PosterPath
	movie.Popularity = movieRequest.Popularity
	movie.Tags = movieRequest.Tags

	newMovie, err := s.repository.Update(movie)
	return newMovie, err
}

func (s *service) Delete(ID int) (Movie, error) {
	movie, err := s.repository.FindById(ID)
	deletedMovie, err := s.repository.Delete(movie)
	return deletedMovie, err
}
