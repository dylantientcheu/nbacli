package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"nbacli/nag/params"
)

// CommonAllPlayers wraps request to and response from commonallplayers endpoint.
type CommonAllPlayers struct {
	*Client
	IsOnlyCurrentSeason int
	LeagueID            string
	Season              string

	Response *Response
}

// NewCommonAllPlayers creates a default CommonAllPlayers instance.
func NewCommonAllPlayers() *CommonAllPlayers {
	return &CommonAllPlayers{
		Client:              NewDefaultClient(),
		IsOnlyCurrentSeason: 0,
		LeagueID:            params.LeagueID.Default(),
		Season:              params.CurrentSeason,
	}
}

// Get sends a GET request to commonallplayers endpoint.
func (c *CommonAllPlayers) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/commonallplayers", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("IsOnlyCurrentSeason", strconv.Itoa(c.IsOnlyCurrentSeason))
	q.Add("LeagueID", c.LeagueID)
	q.Add("Season", c.Season)
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
