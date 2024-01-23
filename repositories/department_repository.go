package repositories

import (
	"github.com/dhimweray222/employee-app/models"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db}
}

func (repo *DepartmentRepository) CreateDepartment(department *models.Department) error {
	result := repo.db.Create(department)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *DepartmentRepository) GetAllDepartments() ([]models.Department, error) {
	var departments []models.Department
	result := repo.db.Find(&departments)
	if result.Error != nil {
		return nil, result.Error
	}
	return departments, nil
}

func (repo *DepartmentRepository) GetDepartmentByID(id uint) (*models.Department, error) {
	var department models.Department
	result := repo.db.First(&department, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &department, nil
}

func (repo *DepartmentRepository) UpdateDepartment(department *models.Department) (*models.Department, error) {
	// Retrieve the existing department from the database
	existingDepartment, err := repo.GetDepartmentByID(department.ID)
	if err != nil {
		return &models.Department{}, err
	}

	// Update the necessary fields
	existingDepartment.DepartmentName = department.DepartmentName
	existingDepartment.MaxClockInTime = department.MaxClockInTime
	existingDepartment.MaxClockOutTime = department.MaxClockOutTime

	// Save the changes
	result := repo.db.Save(existingDepartment)
	if result.Error != nil {
		return &models.Department{}, result.Error
	}
	return existingDepartment, nil
}

func (repo *DepartmentRepository) DeleteDepartment(id uint) error {
	result := repo.db.Delete(&models.Department{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
