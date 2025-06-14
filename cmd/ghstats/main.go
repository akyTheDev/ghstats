package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"

	"github.com/akyTheDev/ghstats/internal/cache"
	"github.com/akyTheDev/ghstats/internal/config"
	"github.com/akyTheDev/ghstats/internal/github"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	repoNamePtr := flag.String("repo", "", "The repository to look up in 'owner/repo' format (e.g., 'google/go-github')")

	flag.Parse()

	if *repoNamePtr == "" {
		logger.Println("Error: The --repo flag is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	repoName := *repoNamePtr

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Config couldn't be loaded: %v", err)
	}

	rc, err := cache.NewRedisClient(ctx, cfg.RedisURL)
	if err != nil {
		logger.Fatalf("Couldn't connect to redis: %v", err)
	}

	gh := github.NewClient()
	stats, err := rc.GetRepoStats(ctx, repoName)

	if err != nil {
		if errors.Is(err, cache.ErrNotFound) {
			stats, err = gh.GetRepoStats(ctx, repoName)
			if err != nil {
				logger.Fatalf("Failed while reading from GitHub api: %v", err)
			}

			err = rc.SetRepoStats(ctx, repoName, stats)
			if err != nil {
				logger.Fatalf("Failed while setting cache: %v", err)
			}
		} else {
			logger.Fatalf("Failed while reading cache from redis: %v", err)
		}
	}

	stats.Display()
}
