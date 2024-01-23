// controllers/department_controller.go

package controllers

import (
	"net/http"
	"strconv"

	"github.com/dhimweray222/employee-app/models"
	"github.com/dhimweray222/employee-app/services"
	"github.com/gofiber/fiber/v2"
)

type DepartmentController struct {
	departmentService *services.DepartmentService
}

func NewDepartmentController(departmentService *services.DepartmentService) *DepartmentController {
	return &DepartmentController{departmentService}
}

func (c *DepartmentController) SetupRoutes(app *fiber.App) {
	// Define routes for the DepartmentController
	departmentGroup := app.Group("/departments")
	{
		departmentGroup.Post("/", c.CreateDepartmentHandler)
		departmentGroup.Get("/", c.GetAllDepartmentsHandler)
		departmentGroup.Get("/:id", c.GetDepartmentByIDHandler)
		departmentGroup.Put("/:id", c.UpdateDepartmentHandler)
		departmentGroup.Delete("/:id", c.DeleteDepartmentHandler)
	}
}

func (c *DepartmentController) CreateDepartmentHandler(ctx *fiber.Ctx) error {
	var department models.Department

	if err := ctx.BodyParser(&department); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := c.departmentService.CreateDepartment(&department)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create department"})
	}

	return ctx.Status(http.StatusCreated).JSON(department)
}

func (c *DepartmentController) GetAllDepartmentsHandler(ctx *fiber.Ctx) error {
	departments, err := c.departmentService.GetAllDepartments()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get departments"})
	}

	return ctx.JSON(departments)
}

func (c *DepartmentController) GetDepartmentByIDHandler(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}

	department, err := c.departmentService.GetDepartmentByID(uint(id))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
	}

	return ctx.JSON(department)
}

func (c *DepartmentController) UpdateDepartmentHandler(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}

	var updatedDepartment models.Department
	if err := ctx.BodyParser(&updatedDepartment); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Set the ID of the department to be updated
	updatedDepartment.ID = uint(id)

	data, err := c.departmentService.UpdateDepartment(&updatedDepartment)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update department"})
	}

	return ctx.Status(http.StatusOK).JSON(data)
}

func (c *DepartmentController) DeleteDepartmentHandler(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}

	err = c.departmentService.DeleteDepartment(uint(id))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete department"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Department deleted successfully"})
}
