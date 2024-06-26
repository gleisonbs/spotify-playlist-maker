package main

import (
	"github.com/gleisonbs/spotify-utils/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("Healthy")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/playlist", controllers.ListPlaylists)
	app.Post("/:userId/playlist/new", controllers.CreatePlaylistWithRecommendations)
	app.Get("/playlist/:playlistId", controllers.GetPlaylistItems)
	app.Get("/recommendations", controllers.ListTrackRecommendations)
	app.Get("/search", controllers.SearchItems)
	app.Get("/health", healthCheck)
	app.Listen(":5555")
}
