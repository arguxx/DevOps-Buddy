package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	  engine := html.New("./views", ".html")

		app := fiber.New(fiber.Config{
			Views: engine, 
    })

    app.Get("/halaman-html", func(c *fiber.Ctx) error {
			data := fiber.Map{
					"Title": "Test passing Static Data",
					"Message": "test long long desc.",
			}
			return c.Render("home", data) 
	  })
    // app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    app.Get("/halaman-lain", func(c *fiber.Ctx) error {
        return c.SendString("kuda-halaman-lain")
    })


    app.Listen(":3000")
}