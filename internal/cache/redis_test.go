package cache

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/akyTheDev/ghstats/internal/models"
	"github.com/go-redis/redismock/v9"
)

const redisURL string = "redis://:redispassword@localhost:6379"

func TestNewRedisClient_Success(t *testing.T) {
	_, err := NewRedisClient(context.Background(), redisURL)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestNewRedisClient_Fail(t *testing.T) {
	tests := []struct {
		name          string
		url           string
		expectedError string
	}{
		{
			name:          "Wrong url",
			url:           "redis",
			expectedError: "Failed while parsing url",
		},
		{
			name:          "Wrong url",
			url:           "redis://:wrongpw@localhost:6379",
			expectedError: "Failed to connect",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewRedisClient(context.Background(), tc.url)

			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			if !strings.Contains(err.Error(), tc.expectedError) {
				t.Fatalf("Expected error string: %s, got %s", tc.expectedError, err.Error())
			}
		})

	}
}

func TestSetRepoStats_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{rdb: db}
	ctx := context.Background()

	stats := &models.Stats{
		ID:   5,
		Name: "TEST REPO",
	}

	expectedStats, err := json.Marshal(stats)
	if err != nil {
		t.Fatal("Error while serializing the stats")
	}

	mock.ExpectSet("user/test-repo", expectedStats, time.Hour).SetVal("OK")

	err = client.SetRepoStats(ctx, "user/test-repo", stats)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if mock.ExpectationsWereMet() != nil {
		t.Fatalf("Called with wrong params")
	}
}

func TestSetRepoStats_Fail(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{rdb: db}
	ctx := context.Background()

	key := "user/test-repo"
	stats := &models.Stats{
		ID:   6,
		Name: "ANOTHER REPO",
	}

	expectedStats, err := json.Marshal(stats)
	if err != nil {
		t.Fatal("Error while serializing the stats")
	}

	expectedErr := errors.New("redis is down")
	mock.ExpectSet(key, expectedStats, time.Hour).SetErr(expectedErr)

	err = client.SetRepoStats(ctx, key, stats)

	if err == nil {
		t.Fatalf("Expected error %v, got nil", err)
	}

	if mock.ExpectationsWereMet() != nil {
		t.Fatal("Called with wrong params")
	}
}

func TestGetRepoStats_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{rdb: db}
	ctx := context.Background()

	key := "user/test-repo"

	expectedStats := &models.Stats{
		Name:            "Test Repo",
		ForksCount:      101,
		StargazersCount: 202,
	}

	// This is the JSON string that we tell the mock Redis to return.
	jsonData, err := json.Marshal(expectedStats)
	if err != nil {
		t.Fatalf("Failed while serializing: %v", err)
	}

	mock.ExpectGet(key).SetVal(string(jsonData))

	gotStats, err := client.GetRepoStats(ctx, key)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(gotStats, expectedStats) {
		t.Errorf("Expected stats: %v, got :%v", expectedStats, gotStats)
	}

	if mock.ExpectationsWereMet() != nil {
		t.Fatal("Called with wrong params")
	}
}

func TestGetRepoStats_FailedNotFound(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{rdb: db}
	ctx := context.Background()

	key := "user/test-repo"

	mock.ExpectGet(key).RedisNil()

	stats, err := client.GetRepoStats(ctx, key)

	if stats != nil {
		t.Fatalf("Expected no stats, got: %v", stats)
	}

	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got: %v", err)
	}
}

func TestGetRepoStats_Failed(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{rdb: db}
	ctx := context.Background()

	key := "user/test-repo"

	expectedErr := errors.New("simulated redis connection error")
	mock.ExpectGet(key).SetErr(expectedErr)

	stats, err := client.GetRepoStats(ctx, key)

	if stats != nil {
		t.Fatalf("Expected no stats, got: %v", stats)
	}

	if !strings.Contains(err.Error(), "Failed to get repo stats from redis") {
		t.Errorf("Expected 'Failed to get repo stats from redis', got: %v", err)
	}
}

func TestGetRepoStats_MarshallFailed(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{rdb: db}
	ctx := context.Background()

	key := "user/test-repo"

	mock.ExpectGet(key).SetVal("BASIC STRING")

	stats, err := client.GetRepoStats(ctx, key)

	if stats != nil {
		t.Fatalf("Expected no stats, got: %v", stats)
	}

	if !strings.Contains(err.Error(), "Failed to unmarshal repo stats from redis") {
		t.Errorf("Expected 'Failed to unmarshal repo stats from redis', got: %v", err)
	}
}
