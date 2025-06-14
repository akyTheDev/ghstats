package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/akyTheDev/ghstats/internal/utils"
)

type Owner struct {
	User string `json:"login"`
}

type Stats struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	RepoOwner       Owner     `json:"owner"`
	HtmlUrl         string    `json:"html_url"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	ForksCount      int       `json:"forks_count"`
	OpenIssuesCount int       `json:"open_issues_count"`
	Language        string    `json:"language"`
}

// Display prints the repository statistics in a formatted, user-friendly table.
func (s *Stats) Display() {
	width := len(s.FullName) + 4

	// --- Header Box ---
	fmt.Println("┌" + strings.Repeat("─", width) + "┐")
	fmt.Printf("│  %s  │\n", s.FullName)
	fmt.Println("└" + strings.Repeat("─", width) + "┘")
	fmt.Println() // Newline for spacing

	if s.Description != "" {
		fmt.Printf("> %s\n\n", s.Description)
	}

	fmt.Println(strings.Repeat("─", width+2))

	statItems := []struct {
		Key   string
		Value string
		Emoji string
	}{
		{"Language:", s.Language, ""},
		{"Stars:", utils.FormatWithCommas(s.StargazersCount), "⭐"},
		{"Forks:", utils.FormatWithCommas(s.ForksCount), "🔱"},
		{"Open Issues:", utils.FormatWithCommas(s.OpenIssuesCount), "❗"},
		{"Last Updated:", s.UpdatedAt.Format("2006-01-02 15:04:05"), ""},
	}

	maxKeyLength := 0
	for _, item := range statItems {
		if len(item.Key) > maxKeyLength {
			maxKeyLength = len(item.Key)
		}
	}

	for _, item := range statItems {
		if item.Value == "" {
			continue
		}
		fmt.Printf("%-*s %s %s\n", maxKeyLength+2, item.Key, item.Value, item.Emoji)
	}
	fmt.Println()
	fmt.Printf("🔗 %s\n", s.HtmlUrl)
}
