package main

import (
	"funnycode/LeagueOfLegend"
	"funnycode/LuckySpin"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Post("/api/v1/reroll", LeagueOfLegend.ReRoll)
	app.Post("/api/v1/lucky-spin", LuckySpin.LuckySpin)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}