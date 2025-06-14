package github

import (
	"net/http"
	"time"
)

type GithubClient struct {
	url    string
	client *http.Client
}

const apiURL = "https://api.github.com"

func NewClient() *GithubClient {
	return &GithubClient{
		url: apiURL,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}
