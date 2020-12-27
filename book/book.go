package book

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	fmt.Println("here")
	c.Send(("All Books"))
}

func GetBook(c *fiber.Ctx) {
	c.Send(("Singular Book"))
}
func AddBook(c *fiber.Ctx) {
	c.Send(("Add a new Book"))
}
func DeleteBook(c *fiber.Ctx) {
	c.Send(("Dellete a Book"))
}
