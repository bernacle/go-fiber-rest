package main

import (
	"fmt"

	"github.com/bernacle/go-fiber-rest/book"
	"github.com/bernacle/go-fiber-rest/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected successfully")

	database.DBConn.AutoMigrate(&book.Book{})

	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)

	app.Listen("localhost:3000")
}
