package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dylantientcheu/nbacli/nag/params"
)

// BoxScoreMiscV2 wraps request to and response from boxscoremiscv2 endpoint.
type BoxScoreMiscV2 struct {
	*Client
	GameID      string
	StartRange  string
	EndRange    string
	RangeType   string
	StartPeriod string
	EndPeriod   string

	Response *Response
}

// NewBoxScoreMiscV2 creates a default BoxScoreMiscV2 instance.
func NewBoxScoreMiscV2(id string) *BoxScoreMiscV2 {
	return &BoxScoreMiscV2{
		Client:      NewDefaultClient(),
		GameID:      id,
		StartRange:  params.DefaultStartRange,
		EndRange:    params.DefaultEndRange,
		RangeType:   params.DefaultRangeType,
		StartPeriod: params.Period.Default(),
		EndPeriod:   params.Period.Default(),
	}
}

// Get sends a GET request to boxscoremiscv2 endpoint.
func (c *BoxScoreMiscV2) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/boxscoremiscv2", c.BaseURL.String()), nil)
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
