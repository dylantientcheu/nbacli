package nag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dylantientcheu/nbacli/nag/params"
)

// GameRotation wraps request to and response from gamerotation endpoint.
type GameRotation struct {
	*Client
	GameID   string
	LeagueID string

	Response *Response
}

// NewGameRotation creates a default GameRotation instance.
func NewGameRotation(id string) *GameRotation {
	return &GameRotation{
		Client:   NewDefaultClient(),
		GameID:   id,
		LeagueID: params.LeagueID.Default(),
	}
}

// Get sends a GET request to gamerotation endpoint.
func (c *GameRotation) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gamerotation", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("LeagueID", c.LeagueID)
	q.Add("GameID", c.GameID)
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
