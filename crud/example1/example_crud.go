package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	// get destination /hello and output "Hello world"
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	// output "Hello world" by port 8080
	app.Listen(":8080")
}
