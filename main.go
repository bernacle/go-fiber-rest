package main

import (
	"github.com/bernacle/go-fiber-rest/book"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(context *fiber.Ctx) error {
	return context.SendString("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Post("/api/v1/books", book.NewBook)
	app.Delete("api/v1/books/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen("localhost:3000")
}
