package dto

import "github.com/gleisonbs/spotify-utils/model"

type Track struct {
	Items []model.Track `json:"items"`
}

type SearchResponse struct {
	Tracks Track `json:"tracks"`
}
