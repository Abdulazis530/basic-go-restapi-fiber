package main

import (
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("<h1>Hello world</h1>")
}

func main() {
	app := fiber.New()
	app.Get("/", helloWorld)

	app.Listen(3000)
}
