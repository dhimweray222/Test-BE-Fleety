package repositories

import (
	"log"

	"github.com/dhimweray222/employee-app/models"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{db}
}

func (repo *AttendanceRepository) CreateAttendance(attendance models.Attendance) (models.Attendance, error) {
	data := models.Attendance{
		AttendanceID: attendance.AttendanceID,
		EmployeeID:   attendance.EmployeeID,
		ClockInTime:  attendance.ClockInTime,
		ClockOutTime: attendance.ClockOutTime,
	}
	result := repo.db.Create(&data)
	if result.Error != nil {
		log.Println("sini")
		return models.Attendance{}, result.Error
	}

	return data, nil
}

func (repo *AttendanceRepository) GetAllAttendances() ([]models.AttendanceWithEmployee, error) {
	var attendances []models.AttendanceWithEmployee
	result := repo.db.
		Table("attendances").
		Select("attendances.*, employees.name as employee_name").
		Joins("INNER JOIN employees ON attendances.employee_id = employees.employee_id").
		Scan(&attendances)
	if result.Error != nil {
		return nil, result.Error
	}
	return attendances, nil
}

func (repo *AttendanceRepository) GetAttendanceByID(attendanceID string) (*models.AttendanceWithEmployee, error) {
	var attendanceWithEmployee models.AttendanceWithEmployee
	result := repo.db.
		Table("attendances").
		Select("attendances.*, employees.name as employee_name").
		Joins("INNER JOIN employees ON attendances.employee_id = employees.employee_id").
		Where("attendances.attendance_id = ?", attendanceID).
		Scan(&attendanceWithEmployee)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &attendanceWithEmployee, nil
}

func (repo *AttendanceRepository) UpdateAttendance(attendance models.Attendance) (models.Attendance, error) {
	var data models.Attendance
	result := repo.db.Where("attendance_id = ? AND employee_id = ?", attendance.AttendanceID, attendance.EmployeeID).First(&data)
	if result.Error != nil {
		return models.Attendance{}, result.Error
	}

	// Update the fields of the found attendance with the new values
	data.ClockOutTime = attendance.ClockOutTime

	// Save the updated attendance
	result = repo.db.Save(&data)
	if result.Error != nil {
		return models.Attendance{}, result.Error
	}
	return data, nil
}

func (repo *AttendanceRepository) FindPreciseAttendanceByDepartment(departmentID int, filter models.QueryFilter) ([]models.AttendanceWithEmployee, error) {
	var preciseAttendances []models.AttendanceWithEmployee
	log.Println(filter)
	log.Println(departmentID)

	query := repo.db.
		Table("attendances").
		Select("attendances.*, employees.name as employee_name").
		Joins("INNER JOIN employees ON attendances.employee_id = employees.id").
		Joins("INNER JOIN departments ON employees.department_id = departments.id").
		Where("departments.id = ?", departmentID)

	// Add clock_in_time and clock_out_time conditions if provided
	if filter.Filter == "clock_in_time" {
		query = query.Where("attendances.clock_in_time <= departments.max_clock_in_time")
	}

	if filter.Filter == "clock_out_time" {
		query = query.Where("attendances.clock_out_time >= departments.max_clock_out_time")
	}

	// Conditionally include date filtering if targetDate is provided
	if !filter.Date.IsZero() {
		query = query.Where("DATE(attendances.created_at) >= ?", filter.Date)
	}

	result := query.Scan(&preciseAttendances)

	if result.Error != nil {
		return nil, result.Error
	}

	return preciseAttendances, nil
}
