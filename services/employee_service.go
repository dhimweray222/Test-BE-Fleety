// services/employee_service.go

package services

import (
	"log"

	"github.com/dhimweray222/employee-app/models"
	"github.com/dhimweray222/employee-app/repositories"
)

type EmployeeService struct {
	employeeRepo *repositories.EmployeeRepository
}

func NewEmployeeService(employeeRepo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepo}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) (models.Employee, error) {

	employee.GenerateID()
	data, err := s.employeeRepo.CreateEmployee(employee)
	if err != nil {
		return models.Employee{}, nil
	}
	// Add any business logic or validation here before creating the employee
	return data, nil
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.employeeRepo.GetAllEmployees()
}

func (s *EmployeeService) GetEmployeeByID(employeeID string) (*models.Employee, error) {
	return s.employeeRepo.GetEmployeeByID(employeeID)
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) (*models.Employee, error) {
	data, err := s.employeeRepo.GetEmployeeByID(employee.EmployeeID)
	if err != nil {
		return &models.Employee{}, err
	}
	log.Println(data)
	data.DepartmentID = employee.DepartmentID
	data.Address = employee.Address
	data.Name = employee.Name
	err = s.employeeRepo.UpdateEmployee(data)
	return data, err
}

func (s *EmployeeService) DeleteEmployee(employeeID string) error {
	return s.employeeRepo.DeleteEmployee(employeeID)
}
