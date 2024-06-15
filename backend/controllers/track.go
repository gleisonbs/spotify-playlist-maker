package controllers

import (
	"encoding/json"
	"log"

	"github.com/gleisonbs/spotify-utils/usecases"
	"github.com/gofiber/fiber/v2"
)

func ListTrackRecommendations(ctx *fiber.Ctx) error {
	trackIds := ctx.Query("trackIds")
	targetPopularity := ctx.Query("popularity")

	headers := ctx.GetReqHeaders()
	authorizationHeader := headers["Authorization"]

	tracks := usecases.GetTrackRecommendations(authorizationHeader[1], trackIds, targetPopularity)

	response, err := json.Marshal(tracks)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.SendString(string(response))
}
