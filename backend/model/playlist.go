package model

type Playlist struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Item struct {
	Track Track `json:"track"`
}
