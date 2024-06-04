package initializers

import (
	"github.com/alifhaider/BCS-Journey-Server/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Badge{})
}
