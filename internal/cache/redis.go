package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/akyTheDev/ghstats/internal/models"
	"github.com/redis/go-redis/v9"
)

var ErrNotFound = errors.New("CACHE: Key not found")

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient(ctx context.Context, redisURL string) (*RedisClient, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("ERROR: NewRedisClient: Failed while parsing url: %v", err)
	}

	rdb := redis.NewClient(opts)

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ERROR: NewRedisClient: Failed to connect: %v", err)
	}

	return &RedisClient{
		rdb: rdb,
	}, nil
}

func (c *RedisClient) SetRepoStats(ctx context.Context, repoName string, stats *models.Stats) error {
	data, err := json.Marshal(stats)
	if err != nil {
		return fmt.Errorf("ERROR: SetRepoStats: Failed while serializing: %v", err)
	}
	return c.rdb.Set(ctx, repoName, data, time.Hour).Err()
}

func (c *RedisClient) GetRepoStats(ctx context.Context, repoName string) (*models.Stats, error) {
	data, err := c.rdb.Get(ctx, repoName).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("ERROR: GetRepoStats: Failed to get repo stats from redis: %w", err)

	}

	var stats models.Stats

	if err := json.Unmarshal([]byte(data), &stats); err != nil {
		return nil, fmt.Errorf("ERROR: GetRepoStats: Failed to unmarshal repo stats from redis: %w", err)
	}

	return &stats, nil
}
