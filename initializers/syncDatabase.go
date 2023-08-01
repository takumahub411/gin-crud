package initializers

import (
	"gin-curd/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Movie{})
}
