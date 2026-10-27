package initializers

import (
	"backend-dev-path/src/models"
)



func MigrateDatabase() { 
	DB.AutoMigrate(
		&models.Users{},
		&models.Courses{},
	)
}