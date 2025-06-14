package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/akyTheDev/ghstats/internal/models"
)

func (gh *GithubClient) GetRepoStats(context context.Context, repoName string) (*models.Stats, error) {
	var result models.Stats

	resp, err := gh.client.Get(fmt.Sprintf("%s/repos/%s", gh.url, repoName))
	if err != nil {
		return nil, fmt.Errorf("ERROR: GetRepoStats: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ERROR: GetRepoStats: Request failed with %d!", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ERROR: GetRepoStats: Read body failed %v", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("ERROR: GetRepoStats: Failed while parsing %v", err)
	}

	return &result, nil
}
