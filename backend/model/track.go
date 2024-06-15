package model

type Track struct {
	Name       string   `json:"name"`
	Id         string   `json:"id"`
	Artist     []Artist `json:"artists"`
	Album      Album    `json:"album"`
	Popularity int      `json:"popularity"`
	Uri        string   `json:"uri"`
}

type TracksToAdd struct {
	Uris     []string `json:"uris"`
	Position int      `json:"position"`
}
