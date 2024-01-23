// models/employee.go

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeID   string `gorm:"type:char(36);primaryKey" json:"employee_id"`
	DepartmentID uint   `gorm:"index" json:"department_id"`
	Department   Department
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (employee *Employee) GenerateID() {
	uuid := uuid.New().String()
	employee.EmployeeID = uuid
}
