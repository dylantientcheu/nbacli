package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nba-cli/nag/params"
)

// TeamGameLog wraps request to and response from teamgamelog endpoint.
type TeamGameLog struct {
	*Client
	LeagueID   string
	TeamID     string
	Season     string
	SeasonType params.SeasonType
	DateFrom   string
	DateTo     string

	Response *Response
}

// NewTeamGameLog creates a default TeamGameLog instance.
func NewTeamGameLog(id string) *TeamGameLog {
	return &TeamGameLog{
		Client:     NewDefaultClient(),
		LeagueID:   params.LeagueID.Default(),
		TeamID:     id,
		Season:     params.CurrentSeason,
		SeasonType: params.AllStar,
	}
}

// Get sends a GET request to teamgamelog endpoint.
func (c *TeamGameLog) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/teamgamelog", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("LeagueID", c.LeagueID)
	q.Add("TeamID", c.TeamID)
	q.Add("Season", c.Season)
	q.Add("SeasonType", string(c.SeasonType))
	q.Add("DateFrom", c.DateFrom)
	q.Add("DateTo", c.DateTo)
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
