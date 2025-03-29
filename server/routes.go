package server

import (
	"github.com/danfelab/youthcongressnepal/routes"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {	

	v1 := app.Group("/v1")
	{
		v1.Post("/register", registerV1)
		v1.Get("/users", getUsersV1)
	}
}
