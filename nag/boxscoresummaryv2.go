package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BoxScoreSummaryV2 wraps request to and response from boxscoresummaryv2 endpoint.
type BoxScoreSummaryV2 struct {
	*Client
	GameID string

	Response *Response
}

// NewBoxScoreSummaryV2 creates a default BoxScoreSummaryV2 instance.
func NewBoxScoreSummaryV2(id string) *BoxScoreSummaryV2 {
	return &BoxScoreSummaryV2{
		Client: NewDefaultClient(),
		GameID: id,
	}
}

// Get sends a GET request to boxscoresummaryv2 endpoint.
func (c *BoxScoreSummaryV2) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/boxscoresummaryv2", c.BaseURL.String()), nil)
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
