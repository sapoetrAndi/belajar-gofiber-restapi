package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sapoetrAndi/belajar-gofiber-restapi/controllers/bookcontroller"
	"github.com/sapoetrAndi/belajar-gofiber-restapi/models"
)

func main() {
	//conect to DB
	models.ConnectDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api := app.Group("/api")
	book := api.Group("/books")
	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":3000")
}
