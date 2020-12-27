package book

import (
	"fmt"

	"github.com/Abdulazis530/basic-go-restapi-fiber/db"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	dBase := db.DBconn
	var books []Book
	dBase.Find(&books)

	err := c.JSON(books)

	if err != nil {
		fmt.Println(err)
	}
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
