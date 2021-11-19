package series

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Series, error)
	FindById(ID int) (Series, error)
	Create(series Series) (Series, error)
	Update(series Series) (Series, error)
	Delete(series Series) (Series, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Series, error) {
	var series []Series

	err := r.db.Find(&series).Error

	return series, err
}

func (r *repository) FindById(ID int) (Series, error) {
	var series Series

	err := r.db.Find(&series, ID).Error

	return series, err
}

func (r *repository) Create(series Series) (Series, error) {
	err := r.db.Create(&series).Error

	return series, err
}

func (r *repository) Update(series Series) (Series, error) {
	err := r.db.Save(&series).Error

	return series, err
}

func (r *repository) Delete(series Series) (Series, error) {
	err := r.db.Delete(&series).Error

	return series, err
}
