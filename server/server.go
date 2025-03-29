package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func StartServer() {
	
	app := fiber.New()

	Routes(app)		
	
	// Start the server
	app.Listen(":8003")
}