package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dylantientcheu/nbacli/nag/params"
)

// PlayByPlayV2 wraps request to and response from playbyplayv2 endpoint.
type PlayByPlayV2 struct {
	*Client
	GameID      string
	StartPeriod string
	EndPeriod   string

	Response *Response
}

// NewPlayByPlayV2 creates a default PlayByPlayV2 instance.
func NewPlayByPlayV2(id string) *PlayByPlayV2 {
	return &PlayByPlayV2{
		Client:      NewDefaultClient(),
		GameID:      id,
		StartPeriod: params.Period.Default(),
		EndPeriod:   params.Period.Default(),
	}
}

// Get sends a GET request to playbyplayv2 endpoint.
func (c *PlayByPlayV2) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/playbyplayv2", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("GameID", c.GameID)
	q.Add("StartPeriod", c.StartPeriod)
	q.Add("EndPeriod", c.EndPeriod)
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
