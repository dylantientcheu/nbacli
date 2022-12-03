package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nbacli/nag/params"
)

// AssistLeaders wraps request to and response from assistleaders endpoint.
type AssistLeaders struct {
	*Client
	LeagueID     string
	PerMode      params.PerMode
	PlayerOrTeam params.PlayerOrTeam
	Season       string
	SeasonType   params.SeasonType

	Response *Response
}

// NewAssistLeaders creates a default AssistLeaders instance.
func NewAssistLeaders() *AssistLeaders {
	return &AssistLeaders{
		Client:       NewDefaultClient(),
		LeagueID:     params.LeagueID.Default(),
		PerMode:      params.DefaultPerMode,
		PlayerOrTeam: params.DefaultPlayerOrTeam,
		Season:       params.CurrentSeason,
		SeasonType:   params.DefaultSeasonType,
	}
}

// Get sends a GET request to assistleaders endpoint.
func (c *AssistLeaders) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/assistleaders", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("LeagueID", c.LeagueID)
	q.Add("PerMode", string(c.PerMode))
	q.Add("PlayerOrTeam", string(c.PlayerOrTeam))
	q.Add("Season", c.Season)
	q.Add("SeasonType", string(c.SeasonType))
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
