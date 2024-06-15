package dto

import "github.com/gleisonbs/spotify-utils/model"

type GetTrackRecommendationsResponse struct {
	Tracks []model.Track `json:"tracks"`
}
