package series

type Service interface {
	FindAll() ([]Series, error)
	FindById(ID int) (Series, error)
	Create(seriesRequest SeriesRequest) (Series, error)
	Update(ID int, seriesRequest SeriesRequest) (Series, error)
	Delete(ID int) (Series, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Series, error) {
	series, err := s.repository.FindAll()
	return series, err
}

func (s *service) FindById(ID int) (Series, error) {
	series, err := s.repository.FindById(ID)
	return series, err
}

func (s *service) Create(seriesRequest SeriesRequest) (Series, error) {
	series := Series{
		Title:      seriesRequest.Title,
		Overview:   seriesRequest.Overview,
		PosterPath: seriesRequest.PosterPath,
		Popularity: seriesRequest.Popularity,
		Tags:       seriesRequest.Tags,
	}

	newSeries, err := s.repository.Create(series)
	return newSeries, err
}

func (s *service) Update(ID int, seriesRequest SeriesRequest) (Series, error) {
	series, err := s.repository.FindById(ID)

	series.Title = seriesRequest.Title
	series.Overview = seriesRequest.Overview
	series.PosterPath = seriesRequest.PosterPath
	series.Popularity = seriesRequest.Popularity
	series.Tags = seriesRequest.Tags

	newSeries, err := s.repository.Update(series)
	return newSeries, err
}

func (s *service) Delete(ID int) (Series, error) {
	series, err := s.repository.FindById(ID)
	deletedseries, err := s.repository.Delete(series)
	return deletedseries, err
}
