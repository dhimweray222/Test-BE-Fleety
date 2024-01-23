package services

import (
	"time"

	"github.com/dhimweray222/employee-app/models"
	"github.com/dhimweray222/employee-app/repositories"
)

type AttendanceService struct {
	attendanceRepo *repositories.AttendanceRepository
}

func NewAttendanceService(attendanceRepo *repositories.AttendanceRepository) *AttendanceService {
	return &AttendanceService{attendanceRepo}
}

func (s *AttendanceService) CreateAttendance(attendance models.AttendanceRequest) (models.Attendance, error) {

	data := models.Attendance{
		EmployeeID:  attendance.EmployeeID,
		ClockInTime: time.Now(),
	}
	data.GenerateID()
	result, err := s.attendanceRepo.CreateAttendance(data)
	if err != nil {
		return models.Attendance{}, err
	}

	return result, nil
}

func (s *AttendanceService) GetAllAttendances() ([]models.AttendanceWithEmployee, error) {
	return s.attendanceRepo.GetAllAttendances()
}

func (s *AttendanceService) GetAttendanceByID(attendanceID string) (*models.AttendanceWithEmployee, error) {
	return s.attendanceRepo.GetAttendanceByID(attendanceID)
}

func (s *AttendanceService) UpdateAttendance(request models.AttendanceUpdateRequest) (models.Attendance, error) {
	currentTime := time.Now()
	data := models.Attendance{
		EmployeeID:   request.EmployeeID,
		AttendanceID: request.AttendanceID,
		ClockOutTime: &currentTime,
	}
	result, err := s.attendanceRepo.UpdateAttendance(data)
	if err != nil {
		return models.Attendance{}, err
	}
	return result, nil
}

func (s *AttendanceService) FindPreciseAttendanceByDepartment(departmentID int, filter models.QueryFilter) ([]models.AttendanceWithEmployee, error) {
	preciseAttendances, err := s.attendanceRepo.FindPreciseAttendanceByDepartment(departmentID, filter)
	if err != nil {
		return nil, err
	}
	return preciseAttendances, nil
}
