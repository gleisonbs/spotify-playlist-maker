package dto

import "github.com/gleisonbs/spotify-utils/model"

type MyPlaylistsResponse struct {
	Items []model.Playlist `json:"items"`
}

type PlaylistResponse struct {
	Items []model.Item `json:"items"`
}

type CreatePlaylistResponse struct {
	Id string `json:"id"`
}

type AddItemsToPlaylistResponse struct {
	Id string `json:"snapshot_id"`
}
