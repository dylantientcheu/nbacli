package nba

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"nba-cli/nag"
)

type BoxScoreSummary struct {
	GameId           string
	GameDate         string
	GameStatus       string
	Gamecode         string
	HomeTeamId       int64
	HomeTeamName     string
	VisitorTeamId    int64
	VisitorTeamName  string
	HomeTeamScore    int
	VisitorTeamScore int
	ArenaName        string
}

type BoxScoreSummaryRepository struct {
}

func (g *BoxScoreSummaryRepository) GetSingleGame() []nag.Stat {
	sbv2 := nag.NewBoxScoreAdvancedV2("0022200248")
	err := sbv2.Get()
	if err != nil {
		panic(err)
	}
	if sbv2.Response == nil {
		panic("no response")
	}

	n := nag.Map(*sbv2.Response)
	var result nag.NewBoxScoreAdvancedResponse
	mapstructure.Decode(n, &result)

	// new games array
	fmt.Printf("%v", result.PlayerStats[6])

	return result.PlayerStats
}
