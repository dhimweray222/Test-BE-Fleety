package services

import (
	"github.com/dhimweray222/employee-app/models"
	"github.com/dhimweray222/employee-app/repositories"
)

type DepartmentService struct {
	departmentRepo *repositories.DepartmentRepository
}

func NewDepartmentService(departmentRepo *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{departmentRepo}
}

func (s *DepartmentService) CreateDepartment(department *models.Department) error {
	// Add any business logic or validation here before creating the department
	return s.departmentRepo.CreateDepartment(department)
}

func (s *DepartmentService) GetAllDepartments() ([]models.Department, error) {
	return s.departmentRepo.GetAllDepartments()
}

func (s *DepartmentService) GetDepartmentByID(id uint) (*models.Department, error) {
	return s.departmentRepo.GetDepartmentByID(id)
}

func (s *DepartmentService) UpdateDepartment(department *models.Department) (*models.Department, error) {
	data, err := s.departmentRepo.UpdateDepartment(department)
	if err != nil {
		return &models.Department{}, nil
	}
	return data, nil
}

func (s *DepartmentService) DeleteDepartment(id uint) error {
	return s.departmentRepo.DeleteDepartment(id)
}
