package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"html/template"
	spotify "main/types"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var API_URL = "https://api.spotify.com/v1"
var AUTH_URL = "https://accounts.spotify.com/api/token"

func authenticate(context context.Context) spotify.TAuthResponse {
	clientId := os.Getenv("CLIENT_ID")
	secret := os.Getenv("SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")

	authToken := base64.StdEncoding.EncodeToString(
		[]byte(clientId + ":" + secret),
	)

	formUrlEncoded := url.Values{}
	formUrlEncoded.Set("grant_type", "refresh_token")
	formUrlEncoded.Set("refresh_token", refreshToken)

	request, _ := http.NewRequestWithContext(
		context,
		http.MethodPost,
		AUTH_URL,
		strings.NewReader(formUrlEncoded.Encode()),
	)

	request.Header.Set("Authorization", "Basic "+authToken)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var accessToken spotify.TAuthResponse

	json.NewDecoder(response.Body).Decode(&accessToken)

	return accessToken
}

func getCurrentMusic(accessToken string, context context.Context) spotify.TMusicPlaying {
	request, _ := http.NewRequest(
		http.MethodGet,
		API_URL+"/me/player/currently-playing",
		nil,
	)

	request.Header.Set("Authorization", "Bearer "+accessToken)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	response, _ := client.Do(request)

	var musicPlaying spotify.TMusicPlaying

	if response.StatusCode != 200 {
		musicPlaying.AlbumCover = template.URL(os.Getenv("DEFAULT_COVER"))

		return musicPlaying
	}

	defer response.Body.Close()

	var currentlyPlaying spotify.TCurrentlyPlaying

	json.NewDecoder(response.Body).Decode(&currentlyPlaying)

	albumImageUrl := currentlyPlaying.Item.Album.Images[0].URL
	musicPlaying.AlbumCover = getAlbumCoverInB64(albumImageUrl)
	musicPlaying.Name = currentlyPlaying.Item.Name

	for index, artist := range currentlyPlaying.Item.Album.Artists {
		if index != 0 {
			musicPlaying.Artists += ", "
		}

		musicPlaying.Artists += artist.Name
	}

	return musicPlaying
}

func getAlbumCoverInB64(imageUrl string) template.URL {
	response, err := http.Get(imageUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytesBuffer := new(bytes.Buffer)
	bytesBuffer.ReadFrom(response.Body)
	imageB64 := base64.StdEncoding.EncodeToString(bytesBuffer.Bytes())

	return template.URL("data:image/jpeg;base64," + imageB64)
}

func formatHTMLTemplate(musicPlaying spotify.TMusicPlaying) (string, error) {
	var pageTemplate *template.Template

	var err error

	if musicPlaying.Name == "" {
		pageTemplate, err = template.
			New("nomusic.html").
			ParseFiles("../../public/nomusic.html")
	} else {
		pageTemplate, err = template.
			New("music.html").
			ParseFiles("../../public/music.html")
	}

	var buf bytes.Buffer

	if err != nil {
		return "", err
	}

	err = pageTemplate.Execute(&buf, musicPlaying)

	if err != nil {
		return "", err
	}

	return buf.String(), err
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	context := context.Background()

	auth := authenticate(context)
	musicPlaying := getCurrentMusic(auth.AccessToken, context)

	htmlPage, err := formatHTMLTemplate(musicPlaying)

	if err != nil {
		panic(err)
	}

	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "image/svg+xml",
		},
		StatusCode: 200,
		Body:       htmlPage,
	}, nil
}

func main() {
	lambda.Start(handler)
}
