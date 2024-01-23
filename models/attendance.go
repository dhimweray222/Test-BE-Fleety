package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	EmployeeID   string     `gorm:"type:char (36);foreignKey" json:"employee_id"`
	AttendanceID string     `gorm:"type:char(36);primaryKey" json:"attendance_id"`
	ClockInTime  time.Time  `json:"clock_in_time" gorm:"type:time"`
	ClockOutTime *time.Time `gorm:"default:null" gorm:"type:time" json:"clock_out_time"`
}

type AttendanceRequest struct {
	gorm.Model
	EmployeeID   string    `gorm:"type:char (36);foreignKey" json:"employee_id"`
	AttendanceID string    `gorm:"type:char(36);primaryKey" json:"attendance_id"`
	ClockInTime  time.Time `json:"clock_in_time"`
}

type AttendanceUpdateRequest struct {
	gorm.Model
	EmployeeID   string `gorm:"type:char (36);foreignKey" json:"employee_id"`
	AttendanceID string `gorm:"type:char(36);primaryKey" json:"attendance_id"`
}

type AttendanceWithEmployee struct {
	Attendance
	EmployeeName string `json:"employee_name"`
}

type QueryFilter struct {
	Filter string    `json:"filter"`
	Date   time.Time `json:"date"`
}

func (attendance *Attendance) GenerateID() {
	uuid := uuid.New().String()
	attendance.AttendanceID = uuid
}
