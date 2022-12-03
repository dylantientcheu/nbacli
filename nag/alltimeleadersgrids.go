package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"nbacli/nag/params"
)

// AllTimeLeadersGrids wraps request to and response from alltimeleadersgrids endpoint.
type AllTimeLeadersGrids struct {
	*Client
	LeagueID   string
	PerMode    params.PerMode
	SeasonType params.SeasonType
	TopX       int

	Response *Response
}

// NewAllTimeLeadersGrids creates a default AllTimeLeadersGrids instance.
func NewAllTimeLeadersGrids() *AllTimeLeadersGrids {
	return &AllTimeLeadersGrids{
		Client:     NewDefaultClient(),
		LeagueID:   params.LeagueID.Default(),
		PerMode:    params.DefaultPerMode,
		SeasonType: params.DefaultSeasonType,
		TopX:       10,
	}
}

// Get sends a GET request to alltimeleadersgrids endpoint.
func (c *AllTimeLeadersGrids) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/alltimeleadersgrids", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("LeagueID", c.LeagueID)
	q.Add("PerMode", string(c.PerMode))
	q.Add("SeasonType", string(c.SeasonType))
	q.Add("TopX", strconv.Itoa(c.TopX))
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
