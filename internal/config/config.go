package config

import (
	"errors"
	"os"
)

type Config struct {
	GithubToken string
	RedisURL    string
}

// Load environment variables of the project.
func Load() (*Config, error) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return nil, errors.New("GITHUB_TOKEN must be provided!")
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		return nil, errors.New("REDIS_URL must be provided!")
	}

	return &Config{
		GithubToken: githubToken,
		RedisURL:    redisURL,
	}, nil
}
