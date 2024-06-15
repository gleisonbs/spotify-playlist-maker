package usecases

import (
	"github.com/gleisonbs/spotify-utils/config"
	"github.com/gleisonbs/spotify-utils/dto"
	"github.com/gleisonbs/spotify-utils/utils"
)

func SearchItems(accessToken string, term string, itemType string, limit string) dto.SearchResponse {
	url := config.API_BASE_URL + "/search?q=" + term + "&type=" + itemType
	if limit != "" {
		url += "&limit=" + limit
	}
	response := utils.MakeHttpRequest(accessToken, "GET", url, nil)
	defer response.Body.Close()

	var searchResponse dto.SearchResponse
	utils.UnmarshalBody(response.Body, &searchResponse)
	return searchResponse
}
