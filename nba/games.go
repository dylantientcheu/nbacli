package nba

import (
	"fmt"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/nleeper/goment"

	"nba-cli/nag"
	"nba-cli/styles"
)

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
	var status = strings.TrimSpace(g.GameStatus)
	if status[len(status)-2:] == "ET" {
		// upcoming game
		gameTime := GetDateTimeFromESTInUTC(status, g.GameDate)
		moment, _ := goment.Unix(gameTime.Unix())
		now, _ := goment.New()

		// show time from now
		desc = fmt.Sprintf("Tip-off %s | %s", moment.From(now), g.ArenaName)
	} else if status == "Final" {
		// upcoming game
		gameDate := GetDateFromString(g.GameDate).Format("2006-01-02")
		desc = fmt.Sprintf("%s - %s | %s", styles.FinalStyle(), gameDate, g.ArenaName)
	} else {
		// live game
		desc = fmt.Sprintf("%s - %s | %s", styles.LiveStyle(), status, g.ArenaName)
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
