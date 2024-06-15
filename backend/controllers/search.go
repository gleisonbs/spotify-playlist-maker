package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gleisonbs/spotify-utils/usecases"
	"github.com/gofiber/fiber/v2"
)

func SearchItems(ctx *fiber.Ctx) error {
	searchTerm := ctx.Query("q")
	itemType := ctx.Query("type")
	limit := ctx.Query("limit")

	headers := ctx.GetReqHeaders()
	authorizationHeader := headers["Authorization"]

	tracks := usecases.SearchItems(authorizationHeader[0], searchTerm, itemType, limit)

	response, err := json.Marshal(tracks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response))
	return ctx.SendString(string(response))
}
