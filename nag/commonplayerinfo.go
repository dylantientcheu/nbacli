package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nbacli/nag/params"
)

// CommonPlayerInfo wraps request to and response from commonplayerinfo endpoint.
type CommonPlayerInfo struct {
	*Client
	PlayerID string
	LeagueID string

	Response *Response
}

// NewCommonPlayerInfo creates a default CommonPlayerInfo instance.
func NewCommonPlayerInfo(id string) *CommonPlayerInfo {
	return &CommonPlayerInfo{
		Client:   NewDefaultClient(),
		PlayerID: id,
		LeagueID: params.LeagueID.Default(),
	}
}

// Get sends a GET request to commonplayerinfo endpoint.
func (c *CommonPlayerInfo) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/commonplayerinfo", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("PlayerID", c.PlayerID)
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
