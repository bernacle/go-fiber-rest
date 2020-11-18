package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(ctx *fiber.Ctx) error {
	return ctx.SendString("All Books")
}
func GetBook(ctx *fiber.Ctx) error {
	return ctx.SendString("A single Book")
}
func NewBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Add Book")
}
func DeleteBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Deletes a Book")
}
