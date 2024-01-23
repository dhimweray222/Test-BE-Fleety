// controllers/employee_controller.go

package controllers

import (
	"net/http"

	"github.com/dhimweray222/employee-app/models"
	"github.com/dhimweray222/employee-app/services"
	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	employeeService *services.EmployeeService
}

func NewEmployeeController(employeeService *services.EmployeeService) *EmployeeController {
	return &EmployeeController{employeeService}
}

func (c *EmployeeController) SetupRoutes(app *fiber.App) {
	// Define routes for the EmployeeController
	employeeGroup := app.Group("/employees")
	{
		employeeGroup.Post("/", c.CreateEmployeeHandler)
		employeeGroup.Get("/", c.GetAllEmployeesHandler)
		employeeGroup.Get("/:employeeID", c.GetEmployeeByIDHandler)
		employeeGroup.Put("/:employeeID", c.UpdateEmployeeHandler)
		employeeGroup.Delete("/:employeeID", c.DeleteEmployeeHandler)
		// Add more routes as needed
	}
}

func (c *EmployeeController) CreateEmployeeHandler(ctx *fiber.Ctx) error {
	var employee models.Employee

	if err := ctx.BodyParser(&employee); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	data, err := c.employeeService.CreateEmployee(&employee)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create employee"})
	}

	return ctx.Status(http.StatusCreated).JSON(data)
}

func (c *EmployeeController) GetAllEmployeesHandler(ctx *fiber.Ctx) error {
	employees, err := c.employeeService.GetAllEmployees()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get employees"})
	}

	return ctx.JSON(employees)
}

func (c *EmployeeController) GetEmployeeByIDHandler(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("employeeID")

	employee, err := c.employeeService.GetEmployeeByID(employeeID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}

	return ctx.JSON(employee)
}

func (c *EmployeeController) UpdateEmployeeHandler(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("employeeID")

	var updatedEmployee models.Employee
	if err := ctx.BodyParser(&updatedEmployee); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Set the employee ID to be updated
	updatedEmployee.EmployeeID = employeeID

	data, err := c.employeeService.UpdateEmployee(&updatedEmployee)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update employee"})
	}

	return ctx.Status(http.StatusOK).JSON(data)
}

func (c *EmployeeController) DeleteEmployeeHandler(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("employeeID")

	err := c.employeeService.DeleteEmployee(employeeID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete employee"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Employee deleted successfully"})
}
