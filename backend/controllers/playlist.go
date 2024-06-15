package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gleisonbs/spotify-utils/usecases"
	"github.com/gleisonbs/spotify-utils/utils"
	"github.com/gofiber/fiber/v2"
)

func ListPlaylists(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	authorizationHeader := headers["Authorization"][0]

	playlists := usecases.GetMyPlaylists(authorizationHeader)
	response, err := json.Marshal(playlists)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.SendString(string(response))
}

func GetPlaylistItems(ctx *fiber.Ctx) error {
	playlistId := ctx.Params("playlistId")

	headers := ctx.GetReqHeaders()
	authorizationHeader := headers["Authorization"][0]

	tracks := usecases.GetPlaylistTracks(authorizationHeader, playlistId)

	response, err := json.Marshal(tracks)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.SendString(string(response))
}

func CreatePlaylistRoute(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	authorizationHeader := headers["Authorization"][0]

	payload := struct {
		PlaylistName string `json:"playlistName"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		utils.LogError(err)
	}

	userId := ctx.Params("userId")

	usecases.CreatePlaylist(authorizationHeader, payload.PlaylistName, userId)

	return ctx.SendString("")
}

func CreatePlaylistWithRecommendations(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	authorizationHeader := headers["Authorization"][0]

	payload := struct {
		PlaylistName string `json:"playlistName"`
		TrackIds     string `json:"trackIds"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		utils.LogError(err)
	}

	userId := ctx.Params("userId")
	fmt.Println("userId", userId)

	tracks := usecases.GetTrackRecommendations(authorizationHeader, payload.TrackIds, "0")

	createPlaylistResponse := usecases.CreatePlaylist(authorizationHeader, payload.PlaylistName, userId)

	usecases.AddTracksToPlaylist(authorizationHeader, createPlaylistResponse.Id, tracks)
	response, err := json.Marshal(tracks)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.SendString(string(response))
}
