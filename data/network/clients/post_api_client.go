package clients

import (
	"net/http"
	"os"
)

type PostApiClient struct {
	http.Client
	BaseUrl string
}

func NewPostApiClient() *PostApiClient {
	baseUrl := os.Getenv("POST_API_URL")
	return &PostApiClient{
		BaseUrl: baseUrl,
	}
}
