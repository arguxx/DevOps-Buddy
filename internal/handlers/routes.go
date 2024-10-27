package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func InformationRoutes(app *fiber.App){
	app.Get("/", func(c *fiber.Ctx) error {
		data := fiber.Map{
				"Title": "Test landing page",
		}
		return c.Render("index", data) 
	})
	app.Get("/content", func(c *fiber.Ctx) error {
		data := fiber.Map{
				"Title": "Test Content Page",
		}
		return c.Render("content", data) 
	})

}

