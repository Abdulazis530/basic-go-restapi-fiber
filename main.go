package main

import (
	"fmt"
	"log"

	"github.com/Abdulazis530/basic-go-restapi-fiber/book"
	"github.com/Abdulazis530/basic-go-restapi-fiber/db"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/", book.AddBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDb() {
	var err error
	db.DBconn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to databse")
	}
	fmt.Println("Db connected")
	db.DBconn.AutoMigrate()
	fmt.Println("Db Migrated")

}
func main() {
	app := fiber.New()
	initDb()
	defer db.DBconn.Close()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
