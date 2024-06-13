package initializers

import (
	"github.com/keshav1411/gofiber/models"
)

func Migartion() {
	
	DB.AutoMigrate(&models.User{})
	//DB.AutoMigrate(&models.Country{})
	//DB.AutoMigrate(&models.Role{})
	//DB.AutoMigrate(&models.Images{})
}