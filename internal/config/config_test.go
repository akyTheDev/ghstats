package config

import (
	"os"
	"testing"
)

const githubToken string = "test-token"
const redisURL string = "redis://:redispassword@localhost:6379"

func TestLoadSucces(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", githubToken)
	t.Setenv("REDIS_URL", redisURL)

	cfg, err := Load()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.GithubToken != githubToken {
		t.Errorf("Expected Github Token: %s, got %s", githubToken, cfg.GithubToken)
	}

	if cfg.RedisURL != redisURL {
		t.Errorf("Expected Redis URL: %s, got %s", redisURL, cfg.RedisURL)
	}
}

func TestLoadFail(t *testing.T) {
	tests := []struct {
		name          string
		setEnv        func()
		missingVarErr string
	}{
		{
			name: "Missing github token",
			setEnv: func() {
				t.Setenv("REDIS_URL", redisURL)
			},
			missingVarErr: "GITHUB_TOKEN must be provided!",
		},
		{
			name: "Missing redis url",
			setEnv: func() {
				t.Setenv("GITHUB_TOKEN", githubToken)
			},
			missingVarErr: "REDIS_URL must be provided!",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Unsetenv("GITHUB_TOKEN")
			os.Unsetenv("REDIS_URL")
			tc.setEnv()

			cfg, err := Load()

			if cfg != nil {
				t.Fatalf("Expected nil config, got: %v", cfg)
			}

			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			if err.Error() != tc.missingVarErr {
				t.Errorf("Expected error: %s, got: %s", tc.missingVarErr, err.Error())
			}
		})
	}
}
