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
	HomeTeamScore    int
	VisitorTeamScore int
	ArenaName        string
}

func (g BoxScoreSummary) Title() string { return g.HomeTeamName + " vs " + g.VisitorTeamName }

// Description the game description to display in a list
func (g BoxScoreSummary) Description() string {
	var desc = ""
	var status = strings.TrimSpace(g.GameStatus)
	if status[len(status)-2:] == "ET" {
		// upcoming game
		gameTime := GetDateTimeFromESTInUTC(status, g.GameDate)
		moment, _ := goment.Unix(gameTime.Unix())
		now, _ := goment.New()

		// show time from now
		desc = fmt.Sprintf("Tip-off %s | %s", moment.From(now), g.ArenaName)
		desc = styles.DescStyle(desc)
	} else if status == "Final" {
		// passed game
		gameDate := GetDateFromString(g.GameDate).Format("2006-01-02")
		desc = fmt.Sprintf("%s  %s | %s", styles.ScoreStyle(g.HomeTeamScore, g.VisitorTeamScore), styles.DescStyle(gameDate), styles.DescStyle(g.ArenaName))
	} else {
		// live game
		desc = fmt.Sprintf("%s %s - %s | %s", styles.LiveStyle(), styles.ScoreStyle(g.HomeTeamScore, g.VisitorTeamScore), styles.DescStyle(status), styles.DescStyle(g.ArenaName))
		desc = styles.DescText.Render(desc)
	}

	return desc
}

// FilterValue choose what field to use for filtering in a Bubbletea list component
func (g BoxScoreSummary) FilterValue() string { return g.HomeTeamName + " vs " + g.VisitorTeamName }

type ScoreboardRepository struct {
}

func (g *BoxScoreSummaryRepository) GetGames(date time.Time) (scbrd []BoxScoreSummary) {
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
	games := make([]BoxScoreSummary, 0, len(result.GameHeader))

	for _, v := range result.GameHeader {
		var game BoxScoreSummary
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

		// get games scores
		for _, s := range result.LineScore {
			if s.TeamID == v.HomeTeamID {
				game.HomeTeamScore = s.Pts
			}
			if s.TeamID == v.VisitorTeamID {
				game.VisitorTeamScore = s.Pts
			}
		}

		games = append(games, game)
	}
	return games
}

func (g *BoxScoreSummaryRepository) GetGameById(gameId string) {
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
