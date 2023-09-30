package repositories

import (
	"api-ayo-absen/internal/app/models"
	"gorm.io/gorm"
)

type WorkingHoursRepositoryInterface interface {
	GetAll() ([]models.WorkingHours, error)
	FindById(Id int) (models.WorkingHours, error)
	Create(workingHours models.WorkingHours) (models.WorkingHours, error)
	Update(workingHours models.WorkingHours) (models.WorkingHours, error)
	Delete(workingHours models.WorkingHours) (bool, error)
}

type workingHourRepository struct {
	db *gorm.DB
}

func NewWorkingHourRespository(db *gorm.DB) *workingHourRepository {
	return &workingHourRepository{db}
}

func (r *workingHourRepository) GetAll() ([]models.WorkingHours, error) {
	var workingHour []models.WorkingHours
	err := r.db.Find(&workingHour).Error

	return workingHour, err
}

func (r *workingHourRepository) FindById(Id int) (models.WorkingHours, error) {
	var workingHour models.WorkingHours
	err := r.db.Find(&workingHour, Id).Error

	return workingHour, err
}

func (r *workingHourRepository) Create(workingHours models.WorkingHours) (models.WorkingHours, error) {
	err := r.db.Create(&workingHours).Error

	return workingHours, err
}

func (r *workingHourRepository) Update(workingHours models.WorkingHours) (models.WorkingHours, error) {
	err := r.db.Save(&workingHours).Error

	return workingHours, err
}

func (r *workingHourRepository) Delete(workingHours models.WorkingHours) (bool, error) {
	status := true
	err := r.db.Delete(&workingHours).Error
	if err != nil {
		status = false
	}
	return status, err
}
