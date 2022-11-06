package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nba-cli/nag/params"
)

// CommonTeamRoster wraps request to and response from commonteamroster endpoint.
type CommonTeamRoster struct {
	*Client
	TeamID   string
	Season   string
	LeagueID string

	Response *Response
}

// NewCommonTeamRoster creates a default CommonTeamRoster instance.
func NewCommonTeamRoster(id string) *CommonTeamRoster {
	return &CommonTeamRoster{
		Client:   NewDefaultClient(),
		TeamID:   id,
		Season:   params.CurrentSeason,
		LeagueID: params.LeagueID.Default(),
	}
}

// Get sends a GET request to commonteamroster endpoint.
func (c *CommonTeamRoster) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/commonteamroster", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("TeamID", c.TeamID)
	q.Add("Season", c.Season)
	q.Add("LeagueID", c.LeagueID)
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
