// models/department.go

package models

import (
	"time"

	"gorm.io/gorm"
)

// Department model
type Department struct {
	gorm.Model
	DepartmentName  string
	MaxClockInTime  time.Time `gorm:"type:time"`
	MaxClockOutTime time.Time `gorm:"type:time"`
}
