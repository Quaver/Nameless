package db

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Swan/Nameless/config"
)

var scoreboardCacheClient = &http.Client{
	Timeout: 10 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

// UpdateScoreboardCache clears a map's scoreboard cache through api-v2.
func UpdateScoreboardCache(_ *Score, m *Map) error {
	endpoint := strings.TrimRight(config.Data.APIBaseUrl, "/") +
		"/v2/private/scoreboards/" + url.PathEscape(m.MD5) + "/cache"
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("create scoreboard cache invalidation request: %w", err)
	}
	req.Header.Set("X-Internal-Secret", config.Data.InternalAPISecret)
	req.Header.Set("Accept", "application/json")

	response, err := scoreboardCacheClient.Do(req)
	if err != nil {
		return fmt.Errorf("invalidate scoreboard cache: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("scoreboard cache invalidation failed with status - %v", response.StatusCode)
	}
	if _, err := io.Copy(io.Discard, response.Body); err != nil {
		return fmt.Errorf("read scoreboard cache invalidation response: %w", err)
	}

	return nil
}
