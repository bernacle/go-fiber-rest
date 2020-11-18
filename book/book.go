package book

import (
	"github.com/bernacle/go-fiber-rest/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(ctx *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return ctx.JSON(books)
}
func GetBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return ctx.JSON(book)
}
func NewBook(ctx *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := ctx.BodyParser(book); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Error parsing body")
	}
	db.Create(&book)
	return ctx.JSON(book)

}
func DeleteBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No book found with given ID")
	}
	db.Delete(&book)
	return ctx.SendString("Book Successfully deleted")
}
