// repositories/employee_repository.go

package repositories

import (
	"github.com/dhimweray222/employee-app/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

func (repo *EmployeeRepository) CreateEmployee(employee *models.Employee) (models.Employee, error) {
	result := repo.db.Create(employee)
	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	var data models.Employee
	dataFind := repo.db.Where("employee_id = ?", employee.EmployeeID).First(&data)
	if dataFind.Error != nil {
		return models.Employee{}, dataFind.Error
	}
	return data, nil
}

func (repo *EmployeeRepository) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	result := repo.db.Preload("Department").Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	return employees, nil
}

func (repo *EmployeeRepository) GetEmployeeByID(employeeID string) (*models.Employee, error) {
	var employee models.Employee
	result := repo.db.Preload("Department").Where("employee_id = ?", employeeID).First(&employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employee, nil
}

func (repo *EmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	result := repo.db.Save(employee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepository) DeleteEmployee(employeeID string) error {
	result := repo.db.Delete(&models.Employee{}, "employee_id = ?", employeeID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
