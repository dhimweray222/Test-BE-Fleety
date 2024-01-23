package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dhimweray222/employee-app/models"
	"github.com/dhimweray222/employee-app/services"
	"github.com/gofiber/fiber/v2"
)

type AttendanceController struct {
	attendanceService *services.AttendanceService
}

func NewAttendanceController(attendanceService *services.AttendanceService) *AttendanceController {
	return &AttendanceController{attendanceService}
}

func (c *AttendanceController) SetupRoutes(app *fiber.App) {
	attendanceGroup := app.Group("/attendances")
	{
		attendanceGroup.Post("/", c.CreateAttendanceHandler)
		attendanceGroup.Get("/", c.GetAllAttendancesHandler)
		attendanceGroup.Get("/:attendanceID", c.GetAttendanceByIDHandler)
		attendanceGroup.Put("/:attendanceID", c.UpdateAttendanceHandler)
		attendanceGroup.Get("/precise-attendances/:departmentID", c.PreciseAttendancesByDepartment)

		// Add more routes as needed
	}
}

func (c *AttendanceController) CreateAttendanceHandler(ctx *fiber.Ctx) error {

	var request models.AttendanceRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}
	result, err := c.attendanceService.CreateAttendance(request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create attendance"})
	}
	return ctx.Status(http.StatusCreated).JSON(result)
}

func (c *AttendanceController) GetAllAttendancesHandler(ctx *fiber.Ctx) error {
	attendances, err := c.attendanceService.GetAllAttendances()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get attendances"})
	}

	return ctx.JSON(attendances)
}

func (c *AttendanceController) GetAttendanceByIDHandler(ctx *fiber.Ctx) error {
	attendanceID := ctx.Params("attendanceID")

	attendance, err := c.attendanceService.GetAttendanceByID(attendanceID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Attendance not found"})
	}

	return ctx.JSON(attendance)
}

func (c *AttendanceController) UpdateAttendanceHandler(ctx *fiber.Ctx) error {
	var request models.AttendanceUpdateRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	result, err := c.attendanceService.UpdateAttendance(request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update attendance"})
	}

	return ctx.Status(http.StatusOK).JSON(result)
}

func (c *AttendanceController) PreciseAttendancesByDepartment(ctx *fiber.Ctx) error {

	// Parse optional date parameter from URL
	dateParam := ctx.Query("date")
	filter := ctx.Query("filter")
	var QueryFilter models.QueryFilter
	var err error
	targetDate, err := time.Parse("2006-01-02", dateParam)
	QueryFilter.Date = targetDate
	if filter != "" {
		QueryFilter.Filter = filter
		log.Println(filter)
	}

	departmentIDStr := ctx.Params("departmentID")
	departmentID, err := strconv.Atoi(departmentIDStr)
	if err != nil {
		return err
	}
	log.Println("departmentID: ", departmentID)

	// Call the service to retrieve precise attendances
	preciseAttendances, err := c.attendanceService.FindPreciseAttendanceByDepartment(departmentID, QueryFilter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve precise attendances"})
	}

	// Return the result as JSON
	return ctx.JSON(preciseAttendances)
}
