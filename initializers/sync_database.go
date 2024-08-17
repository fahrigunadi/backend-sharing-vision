package initializers

import "github.com/fahrigunadi/backend-sharing-vision/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Article{})
}
