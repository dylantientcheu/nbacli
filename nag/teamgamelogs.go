package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dylantientcheu/nbacli/nag/params"
)

// TeamGameLogs wraps request to and response from teamgamelogs endpoint.
type TeamGameLogs struct {
	*Client
	LeagueID       string
	TeamID         string
	OppTeamID      string
	Season         string
	SeasonSegment  params.SeasonSegment
	SeasonType     params.SeasonType
	DateFrom       string
	DateTo         string
	GameSegment    params.GameSegment
	LastNGames     string
	Location       string
	MeasureType    string
	Month          string
	Outcome        params.Outcome
	PORound        string
	PerMode        params.PerMode
	Period         string
	PlayerID       string
	ShotClockRange params.ShotClockRange
	VsConference   params.Conference
	VsDivision     params.Division

	Response *Response
}

// NewTeamGameLogs creates a default TeamGameLogs instance.
func NewTeamGameLogs() *TeamGameLogs {
	return &TeamGameLogs{
		Client:   NewDefaultClient(),
		LeagueID: params.LeagueID.Default(),
	}
}

// Get sends a GET request to teamgamelogs endpoint.
func (c *TeamGameLogs) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/teamgamelogs", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("LeagueID", c.LeagueID)
	q.Add("TeamID", c.TeamID)
	q.Add("OppTeamID", c.OppTeamID)
	q.Add("Season", c.Season)
	q.Add("SeasonSegment", string(c.SeasonSegment))
	q.Add("SeasonType", string(c.SeasonType))
	q.Add("DateFrom", c.DateFrom)
	q.Add("DateTo", c.DateTo)
	q.Add("GameSegment", string(c.GameSegment))
	q.Add("LastNGames", c.LastNGames)
	q.Add("Location", c.Location)
	q.Add("MeasureType", c.MeasureType)
	q.Add("Month", c.Month)
	q.Add("Outcome", string(c.Outcome))
	q.Add("PORound", c.PORound)
	q.Add("PerMode", string(c.PerMode))
	q.Add("Period", c.Period)
	q.Add("PlayerID", c.PlayerID)
	q.Add("ShotClockRange", string(c.ShotClockRange))
	q.Add("VsConference", string(c.VsConference))
	q.Add("VsDivision", string(c.VsDivision))
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
