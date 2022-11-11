package nba

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/nleeper/goment"

	"nba-cli/nag"
)

/* var (
	// Gm the entry repository for the tui
	Gm *nag.BoxScoreSummaryV2
	// Sb the project repository for the tui
	Sb *nag.ScoreBoardV2
) */

type Game struct {
	GameId           string
	GameDate         string
	GameStatus       string
	Gamecode         string
	HomeTeamId       int64
	HomeTeamName     string
	VisitorTeamId    int64
	VisitorTeamName  string
	HomeTeamScore    int64
	VisitorTeamScore int64
	ArenaName        string
}

func (g Game) Title() string { return g.HomeTeamName + " vs " + g.VisitorTeamName }

// Description the game description to display in a list
func (g Game) Description() string {
	var desc = ""
	if g.GameStatus != "Final" {
		gameTime := GetDateTimeFromESTInUTC(g.GameStatus, g.GameDate)
		moment, _ := goment.Unix(gameTime.Unix())
		now, _ := goment.New()

		desc = fmt.Sprintf("Tip-off %s | %s", moment.From(now), g.ArenaName)
	} else {
		desc = fmt.Sprintf("%s | %s", g.GameDate, g.ArenaName)
	}

	return desc
}

// FilterValue choose what field to use for filtering in a Bubbletea list component
func (g Game) FilterValue() string { return g.HomeTeamName + " vs " + g.VisitorTeamName }

type Repository interface {
	GetGames(date time.Time) (scbrd []Game)
	// GetGameById(gameId string) (scbrd Game)
	// CreateProject(name string) (Project, error)
	// DeleteProject(projectID uint) error
	// RenameProject(projectID uint) error
}

type ScoreboardRepository struct {
}

func (g *ScoreboardRepository) GetGames(date time.Time) (scbrd []Game) {
	sbv2 := nag.NewScoreBoardV2(date)
	err := sbv2.Get()
	if err != nil {
		panic(err)
	}
	if sbv2.Response == nil {
		panic("no response")
	}

	n := nag.Map(*sbv2.Response)
	var result nag.ScoreBoardResponse
	mapstructure.Decode(n, &result)

	// new games array
	games := make([]Game, 0, len(result.GameHeader))

	for _, v := range result.GameHeader {
		var game Game
		game.GameId = v.GameID
		game.GameDate = v.GameDateEst
		game.GameStatus = v.GameStatusText
		game.HomeTeamId = v.HomeTeamID
		game.VisitorTeamId = v.VisitorTeamID
		game.Gamecode = v.Gamecode

		// get team name by id
		hteam, herr := GetTeamByIdOrTricode(v.HomeTeamID, "")
		ateam, aerr := GetTeamByIdOrTricode(v.VisitorTeamID, "")
		if herr != nil {
			panic(herr)
		}
		if aerr != nil {
			panic(aerr)
		}

		game.HomeTeamName = hteam.FullName
		game.VisitorTeamName = ateam.FullName
		game.ArenaName = v.ArenaName
		game.GameStatus = v.GameStatusText

		games = append(games, game)

		/* p, _ := json.MarshalIndent(game, "", " ")
		fmt.Println(string(p)) */
	}
	return games
}

func (g *ScoreboardRepository) GetGameById(gameId string) {
	sbv2 := nag.NewBoxScoreAdvancedV2(gameId)
	err := sbv2.Get()

	if err != nil {
		panic(err)
	}

	if sbv2.Response == nil {
		panic("nil response")
	}

	// m := nag.Map(*sbv2.Response)
	n := nag.Map(*sbv2.Response)
	h := n["GameHeader"].([]map[string]interface{})

	for _, v := range h {
		fmt.Println(v["GAME_ID"])
	}

	// return n
}
