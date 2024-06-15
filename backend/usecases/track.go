package usecases

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gleisonbs/spotify-utils/config"
	"github.com/gleisonbs/spotify-utils/dto"
	"github.com/gleisonbs/spotify-utils/model"
	"github.com/gleisonbs/spotify-utils/utils"
)

func GetTrackRecommendations(accessToken string, trackIds string, targetPopularity string) []model.Track {
	url := config.API_BASE_URL + "/recommendations?seed_tracks=" + trackIds + "&limit=50"
	if targetPopularity != "" {
		url += "&target_popularity=" + targetPopularity
	}
	request, err := http.NewRequest("GET", url, nil)
	utils.LogError(err)

	request.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	utils.LogError(err)

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	utils.LogError(err)

	var getRecommendationsResponse dto.GetTrackRecommendationsResponse
	err = json.Unmarshal(bodyStr, &getRecommendationsResponse)
	utils.LogError(err)

	var tracks []model.Track
	for _, track := range getRecommendationsResponse.Tracks {
		tracks = append(tracks, track)
	}

	return tracks
}
