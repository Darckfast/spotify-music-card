package spotify

import "html/template"

type TMusicPlaying struct {
	Name       string       `json:"musicName"`
	Artists    string       `json:"artists"`
	AlbumCover template.URL `json:"albumCover"`
}
