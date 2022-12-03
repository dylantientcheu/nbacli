package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nbacli/nag/params"
)

// CommonPlayoffSeries wraps request to and response from commonplayoffseries endpoint.
type CommonPlayoffSeries struct {
	*Client
	LeagueID string
	Season   string
	SeriesID string

	Response *Response
}

// NewCommonPlayoffSeries creates a default CommonPlayoffSeries instance.
func NewCommonPlayoffSeries() *CommonPlayoffSeries {
	return &CommonPlayoffSeries{
		Client:   NewDefaultClient(),
		LeagueID: params.LeagueID.Default(),
		Season:   params.CurrentSeason,
	}
}

// Get sends a GET request to commonplayoffseries endpoint.
func (c *CommonPlayoffSeries) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/commonplayoffseries", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("LeagueID", c.LeagueID)
	q.Add("Season", c.Season)
	q.Add("SeriesID", c.SeriesID)
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
