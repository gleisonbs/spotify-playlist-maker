package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gleisonbs/spotify-utils/config"
	"github.com/gleisonbs/spotify-utils/dto"
	"github.com/gleisonbs/spotify-utils/utils"
	"github.com/gofiber/fiber/v2"
)

func Authorize(ctx *fiber.Ctx) error {
	scope := "user-read-private user-read-email playlist-read-private playlist-modify-public playlist-modify-private"
	state := utils.RandSeq(16)

	return ctx.Redirect("https://accounts.spotify.com/authorize?" +
		"response_type=code" +
		"&client_id=" + config.CLIENT_ID +
		"&scope=" + scope +
		"&redirect_uri=" + config.REDIRECT_URI +
		"&state=" + state)
}

func AuthWithCode(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	state := ctx.Query("state")

	formData := url.Values{
		"code":         {code},
		"redirect_uri": {config.REDIRECT_URI},
		"grant_type":   {"authorization_code"},
	}
	encodedFormData := formData.Encode()

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(encodedFormData))
	if err != nil {
		log.Fatalln(err)
	}

	auth := config.CLIENT_ID + ":" + config.CLIENT_SECRET
	auth = base64.StdEncoding.EncodeToString([]byte(auth))

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	response, err := client.Do(request)

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var authResponse dto.AuthResponse
	err = json.Unmarshal(bodyStr, &authResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// config.AccessToken = authResponse.AccessToken

	defer response.Body.Close()
	return ctx.SendString(code + "\n" + state)
}
