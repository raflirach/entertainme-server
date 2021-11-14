package movie

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Movie, error)
	FindById(ID int) (Movie, error)
	Create(movie Movie) (Movie, error)
	Update(movie Movie) (Movie, error)
	Delete(movie Movie) (Movie, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Movie, error) {
	var movies []Movie

	err := r.db.Find(&movies).Error

	return movies, err
}

func (r *repository) FindById(ID int) (Movie, error) {
	var movie Movie

	err := r.db.Find(&movie, ID).Error

	return movie, err
}

func (r *repository) Create(movie Movie) (Movie, error) {
	err := r.db.Create(&movie).Error

	return movie, err
}

func (r *repository) Update(movie Movie) (Movie, error) {
	err := r.db.Save(&movie).Error

	return movie, err
}

func (r *repository) Delete(movie Movie) (Movie, error) {
	err := r.db.Delete(&movie).Error

	return movie, err
}
