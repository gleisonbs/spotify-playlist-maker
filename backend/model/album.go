package model

type Image struct {
	Url string `json:"url"`
}

type Album struct {
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}
