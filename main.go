package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/keshav1411/gofiber/controllers"
	"github.com/keshav1411/gofiber/initializers"
)

func init() {
	initializers.LoadVariable()
	initializers.ConnectToDB()
	// initializers.Migartion()
}

func main() {
	fmt.Println("Run Service")

	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // You can restrict this to specific origins
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Define your routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to golang fiber with gorm",
		})
	})

	app.Post("/api/users", controllers.SignUp)
	app.Get("/api/users", controllers.GetUsers)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Put("/api/updateuser/:id", controllers.UpdateUser)
	app.Delete("/api/deleteuser/:id", controllers.DeleteUser)
	app.Post("/api/login", controllers.Login)

	log.Fatal(app.Listen(":4001"))
}
