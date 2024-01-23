// main.go

package main

import (
	"github.com/dhimweray222/employee-app/config"
	"github.com/dhimweray222/employee-app/controllers"
	"github.com/dhimweray222/employee-app/repositories"
	"github.com/dhimweray222/employee-app/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	departmentRepo := repositories.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepo)
	departmentController := controllers.NewDepartmentController(departmentService)

	employeeRepo := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	attendanceRepo := repositories.NewAttendanceRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	attendanceController := controllers.NewAttendanceController(attendanceService)

	// Set up Fiber app
	app := fiber.New()

	// Use the SetupRoutes method to define routes for the controller
	departmentController.SetupRoutes(app)
	employeeController.SetupRoutes(app)
	attendanceController.SetupRoutes(app)

	// Run the application
	err = app.Listen(":3000")
	if err != nil {
		panic("Failed to start Fiber server: " + err.Error())
	}
}
