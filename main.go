package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		hostname, _ := os.Hostname()
		response := map[string]string{
			"version":   "V4",
			"path":      c.Path(),
			"client_ip": c.IP(),
			"hostname":  hostname,
		}
		return c.JSON(response)
	})

	app.Listen(":8080")
}
