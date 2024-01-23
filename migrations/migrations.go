// migrations/migrate.go

package migrations

import (
	"fmt"

	"github.com/dhimweray222/employee-app/models"
	"gorm.io/gorm"
)

// RunMigrations runs database migrations
func RunMigrations(db *gorm.DB) {
	// AutoMigrate will create the 'departments' table and apply any necessary changes
	db.AutoMigrate(&models.Department{})
	db.AutoMigrate(&models.Employee{})
	db.AutoMigrate(&models.Attendance{})

	fmt.Println("Migrations completed successfully")
}
