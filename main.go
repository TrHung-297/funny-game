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
	app.Put("/api/v1/pick-character", LeagueOfLegend.Pick)
	app.Post("/api/v1/lucky-spin", LuckySpin.LuckySpin)
}

func main() {
	app := fiber.New()
	LeagueOfLegend.Init()
	setupRoutes(app)
	app.Listen(3000)
}