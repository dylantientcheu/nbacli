package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BoxScoreDefensive wraps request to and response from boxscoredefensive endpoint.
type BoxScoreDefensive struct {
	*Client
	GameID string

	Response *Response
}

// NewBoxScoreDefensive creates a default BoxScoreDefensive instance.
func NewBoxScoreDefensive(id string) *BoxScoreDefensive {
	return &BoxScoreDefensive{
		Client: NewDefaultClient(),
		GameID: id,
	}
}

// Get sends a GET request to boxscoredefensive endpoint.
func (c *BoxScoreDefensive) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/boxscoredefensive", c.BaseURL.String()), nil)
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
