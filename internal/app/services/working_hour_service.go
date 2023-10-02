package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
	"errors"
)

type WorkingHoursServiceInterface interface {
	GetAll() ([]models.WorkingHours, error)
	FindById(Id int) (models.WorkingHours, error)
	Create(workingHours request.WorkingHoursRequest) (models.WorkingHours, error)
	Update(Id int, workingHours request.WorkingHoursUpdateRequest) (models.WorkingHours, error)
	Delete(Id int) (bool, error)
}

type workingHourService struct {
	workingHourRepository repositories.WorkingHoursRepositoryInterface
}

func NewWorkingHourService(workingHoursRepositoryInterface repositories.WorkingHoursRepositoryInterface) *workingHourService {
	return &workingHourService{workingHoursRepositoryInterface}
}

func (s *workingHourService) GetAll() ([]models.WorkingHours, error) {
	result, err := s.workingHourRepository.GetAll()

	return result, err
}

func (s *workingHourService) FindById(Id int) (models.WorkingHours, error) {
	result, err := s.workingHourRepository.FindById(Id)

	return result, err
}

func (s *workingHourService) Create(workingHours request.WorkingHoursRequest) (models.WorkingHours, error) {
	workingHour := models.WorkingHours{
		CompanyId: workingHours.CompanyId,
		StartDay:  workingHours.StartDay,
		EndDay:    workingHours.EndDay,
		StartTime: workingHours.StartTime,
		EndTime:   workingHours.EndTime,
		Active:    workingHours.Active,
		CreatedAt: workingHours.CreatedAt,
		CreatedBy: workingHours.CreatedBy,
		UpdatedAt: workingHours.CreatedAt,
	}
	result, err := s.workingHourRepository.Create(workingHour)

	return result, err
}

func (s *workingHourService) Update(Id int, workingHours request.WorkingHoursUpdateRequest) (models.WorkingHours, error) {
	findWorkingHour, err := s.workingHourRepository.FindById(Id)
	if err != nil {
		return models.WorkingHours{}, errors.New("Data tidak ditemukan")
	}
	workingHour := models.WorkingHours{
		Id:        findWorkingHour.Id,
		CompanyId: workingHours.CompanyId,
		StartDay:  workingHours.StartDay,
		EndDay:    workingHours.EndDay,
		StartTime: workingHours.StartTime,
		EndTime:   workingHours.EndTime,
		Active:    workingHours.Active,
		UpdatedAt: workingHours.UpdatedAt,
		UpdatedBy: workingHours.UpdatedBy,
	}
	return s.workingHourRepository.Update(workingHour)
}

func (s *workingHourService) Delete(Id int) (bool, error) {
	findWorkingHour, err := s.workingHourRepository.FindById(Id)
	if err != nil {
		return false, errors.New("Data tidak ditemukan")
	}
	return s.workingHourRepository.Delete(findWorkingHour)
}
