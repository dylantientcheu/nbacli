package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BoxScorePlayerTrackV2 wraps request to and response from boxscoreplayertrackv2 endpoint.
type BoxScorePlayerTrackV2 struct {
	*Client
	GameID string

	Response *Response
}

// NewBoxScorePlayerTrackV2 creates a default BoxScorePlayerTrackV2 instance.
func NewBoxScorePlayerTrackV2(id string) *BoxScorePlayerTrackV2 {
	return &BoxScorePlayerTrackV2{
		Client: NewDefaultClient(),
		GameID: id,
	}
}

// Get sends a GET request to boxscoreplayertrackv2 endpoint.
func (c *BoxScorePlayerTrackV2) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/boxscoreplayertrackv2", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("GameID", c.GameID)
	req.URL.RawQuery = q.Encode()

	b, err := c.Do(req)
	if err != nil {
		return err
	}

	var res Response
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	c.Response = &res
	return nil
}
