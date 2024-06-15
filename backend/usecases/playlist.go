package usecases

import (
	"bytes"
	"encoding/json"

	"github.com/gleisonbs/spotify-utils/config"
	"github.com/gleisonbs/spotify-utils/dto"
	"github.com/gleisonbs/spotify-utils/model"
	"github.com/gleisonbs/spotify-utils/utils"
)

func GetMyPlaylists(accessToken string) []model.Playlist {
	url := config.API_BASE_URL + "/me/playlists?offset=0&limit=50"
	response := utils.MakeHttpRequest(accessToken, "GET", url, nil)
	defer response.Body.Close()

	var myPlaylistsResponse dto.MyPlaylistsResponse
	utils.UnmarshalBody(response.Body, &myPlaylistsResponse)

	return myPlaylistsResponse.Items
}

func GetPlaylistTracks(accessToken string, plalistId string) []model.Track {
	url := config.API_BASE_URL + "/playlists/" + plalistId + "/tracks?offset=0&limit=50"
	response := utils.MakeHttpRequest(accessToken, "GET", url, nil)
	defer response.Body.Close()

	var playlistResponse dto.PlaylistResponse
	utils.UnmarshalBody(response.Body, &playlistResponse)

	var tracks []model.Track
	for _, item := range playlistResponse.Items {
		tracks = append(tracks, item.Track)
	}

	return tracks
}

func AddTracksToPlaylist(accessToken string, plalistId string, tracks []model.Track) dto.AddItemsToPlaylistResponse {
	payload := model.TracksToAdd{
		Uris:     make([]string, 0),
		Position: 0,
	}
	for _, track := range tracks {
		payload.Uris = append(payload.Uris, track.Uri)
	}

	body, err := json.Marshal(payload)
	utils.LogError(err)

	url := config.API_BASE_URL + "/playlists/" + plalistId + "/tracks"
	response := utils.MakeHttpRequest(accessToken, "POST", url, bytes.NewBuffer(body))
	defer response.Body.Close()

	var addItemsToPlaylistResponse dto.AddItemsToPlaylistResponse
	utils.UnmarshalBody(response.Body, &addItemsToPlaylistResponse)
	return addItemsToPlaylistResponse
}

func CreatePlaylist(accessToken string, playlistName string, userId string) dto.CreatePlaylistResponse {
	body, err := json.Marshal(map[string]interface{}{
		"name":        playlistName,
		"description": "API PLAYLIST",
		"public":      false,
	})
	utils.LogError(err)

	response := utils.MakeHttpRequest(accessToken, "POST", config.API_BASE_URL+"/users/"+userId+"/playlists", bytes.NewBuffer(body))
	defer response.Body.Close()

	var createPlaylistResponse dto.CreatePlaylistResponse
	utils.UnmarshalBody(response.Body, &createPlaylistResponse)
	return createPlaylistResponse
}
