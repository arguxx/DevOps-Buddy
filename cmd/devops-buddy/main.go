package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"devops-buddy/internal/handlers"
)

func main() {
	  webInformation := html.New("./web/templates/", ".html")

		app := fiber.New(fiber.Config{
			Views: webInformation, 
    })
		handlers.InformationRoutes(app)

    app.Static("/static", "./web/static")
    app.Listen(":3000")
}