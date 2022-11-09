package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nba-cli/nag/params"
)

// BoxScoreTraditionalV2 wraps request to and response from boxscoretraditionalv2 endpoint.
type BoxScoreTraditionalV2 struct {
	*Client
	GameID      string
	StartRange  string
	EndRange    string
	RangeType   string
	StartPeriod string
	EndPeriod   string

	Response *Response
}

// NewBoxScoreTraditionalV2 creates a default BoxScoreTraditionalV2 instance.
func NewBoxScoreTraditionalV2(id string) *BoxScoreTraditionalV2 {
	return &BoxScoreTraditionalV2{
		Client:      NewDefaultClient(),
		GameID:      id,
		StartRange:  params.DefaultStartRange,
		EndRange:    params.DefaultEndRange,
		RangeType:   params.DefaultRangeType,
		StartPeriod: params.Period.Default(),
		EndPeriod:   params.Period.Default(),
	}
}

// Get sends a GET request to boxscoretraditionalv2 endpoint.
func (c *BoxScoreTraditionalV2) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/boxscoretraditionalv2", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("GameID", c.GameID)
	q.Add("StartRange", c.StartRange)
	q.Add("EndRange", c.EndRange)
	q.Add("RangeType", c.RangeType)
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
