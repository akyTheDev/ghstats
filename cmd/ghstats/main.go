package main

import (
	"context"
	"log"
	"os"

	"github.com/akyTheDev/ghstats/internal/cache"
	"github.com/akyTheDev/ghstats/internal/config"
	"github.com/akyTheDev/ghstats/internal/github"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Config couldn't be loaded: %v", err)
	}

	gh := github.NewClient()

	stats, _ := gh.GetRepoStats(context.Background(), "akyTheDev/currency-bot")

	rc, err := cache.NewRedisClient(ctx, cfg.RedisURL)
	if err != nil {
		logger.Fatalf("Couldn't connect to redis: %v", err)
	}

	err = rc.SetRepoStats(ctx, "akythedev/currency-bot", stats)
	if err != nil {
		logger.Fatalf("Couldn't insert to redis :%v", err)
	}
}
