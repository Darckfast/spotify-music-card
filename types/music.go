package spotify

type TMusicPlaying struct {
	Name       string `json:"musicName"`
	Artists    string `json:"artists"`
	AlbumCover string `json:"albumCover"`
}
