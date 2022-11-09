package nag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"nba-cli/nag/params"
)

// ScoreBoardV2 wraps request to and response from scoreboardv2 endpoint.
type ScoreBoardV2 struct {
	*Client
	DayOffset int
	GameDate  string
	LeagueID  string

	Response *Response
}

type ScoreBoardResponse struct {
	Available              []Available            `mapstructure:"Available"`
	EastConfStandingsByDay []StConfStandingsByDay `mapstructure:"EastConfStandingsByDay"`
	GameHeader             []GameHeader           `mapstructure:"GameHeader"`
	LastMeeting            []LastMeeting          `mapstructure:"LastMeeting"`
	LineScore              []LineScore            `mapstructure:"LineScore"`
	SeriesStandings        interface{}            `mapstructure:"SeriesStandings"`
	TeamLeaders            interface{}            `mapstructure:"TeamLeaders"`
	TicketLinks            interface{}            `mapstructure:"TicketLinks"`
	WestConfStandingsByDay []StConfStandingsByDay `mapstructure:"WestConfStandingsByDay"`
	WinProbability         interface{}            `mapstructure:"WinProbability"`
}

type Available struct {
	GameID      string `mapstructure:"GAME_ID"`
	PtAvailable int64  `mapstructure:"PT_AVAILABLE"`
}

type StConfStandingsByDay struct {
	Conference    Conference    `mapstructure:"CONFERENCE"`
	G             int64         `mapstructure:"G"`
	HomeRecord    string        `mapstructure:"HOME_RECORD"`
	L             int64         `mapstructure:"L"`
	LeagueID      string        `mapstructure:"LEAGUE_ID"`
	RoadRecord    string        `mapstructure:"ROAD_RECORD"`
	SeasonID      string        `mapstructure:"SEASON_ID"`
	Standingsdate Standingsdate `mapstructure:"STANDINGSDATE"`
	Team          string        `mapstructure:"TEAM"`
	TeamID        int64         `mapstructure:"TEAM_ID"`
	W             int64         `mapstructure:"W"`
	WPct          float64       `mapstructure:"W_PCT"`
}

type GameHeader struct {
	ArenaName                     string      `mapstructure:"ARENA_NAME"`
	AwayTvBroadcasterAbbreviation string      `mapstructure:"AWAY_TV_BROADCASTER_ABBREVIATION"`
	Gamecode                      string      `mapstructure:"GAMECODE"`
	GameDateEst                   string      `mapstructure:"GAME_DATE_EST"`
	GameID                        string      `mapstructure:"GAME_ID"`
	GameSequence                  int64       `mapstructure:"GAME_SEQUENCE"`
	GameStatusID                  int64       `mapstructure:"GAME_STATUS_ID"`
	GameStatusText                string      `mapstructure:"GAME_STATUS_TEXT"`
	HomeTeamID                    int64       `mapstructure:"HOME_TEAM_ID"`
	HomeTvBroadcasterAbbreviation string      `mapstructure:"HOME_TV_BROADCASTER_ABBREVIATION"`
	LivePCTime                    string      `mapstructure:"LIVE_PC_TIME"`
	LivePeriod                    int64       `mapstructure:"LIVE_PERIOD"`
	LivePeriodTimeBcast           string      `mapstructure:"LIVE_PERIOD_TIME_BCAST"`
	NatlTvBroadcasterAbbreviation interface{} `mapstructure:"NATL_TV_BROADCASTER_ABBREVIATION"`
	Season                        string      `mapstructure:"SEASON"`
	VisitorTeamID                 int64       `mapstructure:"VISITOR_TEAM_ID"`
	WhStatus                      int64       `mapstructure:"WH_STATUS"`
	WnbaCommissionerFlag          int64       `mapstructure:"WNBA_COMMISSIONER_FLAG"`
}

type LastMeeting struct {
	GameID                       string `mapstructure:"GAME_ID"`
	LastGameDateEst              string `mapstructure:"LAST_GAME_DATE_EST"`
	LastGameHomeTeamAbbreviation string `mapstructure:"LAST_GAME_HOME_TEAM_ABBREVIATION"`
	LastGameHomeTeamCity         string `mapstructure:"LAST_GAME_HOME_TEAM_CITY"`
	LastGameHomeTeamID           int64  `mapstructure:"LAST_GAME_HOME_TEAM_ID"`
	LastGameHomeTeamName         string `mapstructure:"LAST_GAME_HOME_TEAM_NAME"`
	LastGameHomeTeamPoints       int64  `mapstructure:"LAST_GAME_HOME_TEAM_POINTS"`
	LastGameID                   string `mapstructure:"LAST_GAME_ID"`
	LastGameVisitorTeamCity      string `mapstructure:"LAST_GAME_VISITOR_TEAM_CITY"`
	LastGameVisitorTeamCity1     string `mapstructure:"LAST_GAME_VISITOR_TEAM_CITY1"`
	LastGameVisitorTeamID        int64  `mapstructure:"LAST_GAME_VISITOR_TEAM_ID"`
	LastGameVisitorTeamName      string `mapstructure:"LAST_GAME_VISITOR_TEAM_NAME"`
	LastGameVisitorTeamPoints    int64  `mapstructure:"LAST_GAME_VISITOR_TEAM_POINTS"`
}

type LineScore struct {
	AST              interface{} `mapstructure:"AST"`
	Fg3Pct           interface{} `mapstructure:"FG3_PCT"`
	FgPct            interface{} `mapstructure:"FG_PCT"`
	FtPct            interface{} `mapstructure:"FT_PCT"`
	GameDateEst      string      `mapstructure:"GAME_DATE_EST"`
	GameID           string      `mapstructure:"GAME_ID"`
	GameSequence     int64       `mapstructure:"GAME_SEQUENCE"`
	Pts              interface{} `mapstructure:"PTS"`
	PtsOt1           interface{} `mapstructure:"PTS_OT1"`
	PtsOt10          interface{} `mapstructure:"PTS_OT10"`
	PtsOt2           interface{} `mapstructure:"PTS_OT2"`
	PtsOt3           interface{} `mapstructure:"PTS_OT3"`
	PtsOt4           interface{} `mapstructure:"PTS_OT4"`
	PtsOt5           interface{} `mapstructure:"PTS_OT5"`
	PtsOt6           interface{} `mapstructure:"PTS_OT6"`
	PtsOt7           interface{} `mapstructure:"PTS_OT7"`
	PtsOt8           interface{} `mapstructure:"PTS_OT8"`
	PtsOt9           interface{} `mapstructure:"PTS_OT9"`
	PtsQtr1          interface{} `mapstructure:"PTS_QTR1"`
	PtsQtr2          interface{} `mapstructure:"PTS_QTR2"`
	PtsQtr3          interface{} `mapstructure:"PTS_QTR3"`
	PtsQtr4          interface{} `mapstructure:"PTS_QTR4"`
	Reb              interface{} `mapstructure:"REB"`
	TeamAbbreviation string      `mapstructure:"TEAM_ABBREVIATION"`
	TeamCityName     string      `mapstructure:"TEAM_CITY_NAME"`
	TeamID           int64       `mapstructure:"TEAM_ID"`
	TeamName         string      `mapstructure:"TEAM_NAME"`
	TeamWINSLosses   string      `mapstructure:"TEAM_WINS_LOSSES"`
	Tov              interface{} `mapstructure:"TOV"`
}

type Conference string

const (
	East Conference = "East"
	West Conference = "West"
)

type Standingsdate string

const (
	The11062022 Standingsdate = "11/06/2022"
)

// NewScoreBoardV2 creates a default ScoreBoardV2 instance.
func NewScoreBoardV2(date time.Time) *ScoreBoardV2 {
	fmt.Println(date.Format("2006-01-02"))
	return &ScoreBoardV2{
		Client:    NewDefaultClient(),
		DayOffset: 0,
		GameDate:  date.Format("2006-01-02"),
		LeagueID:  params.LeagueID.Default(),
	}
}

// Get sends a GET request to scoreboardv2 endpoint.
func (c *ScoreBoardV2) Get() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/scoreboardv2", c.BaseURL.String()), nil)
	if err != nil {
		return err
	}

	req.Header = DefaultStatsHeader

	q := req.URL.Query()
	q.Add("DayOffset", strconv.Itoa(c.DayOffset))
	q.Add("GameDate", c.GameDate)
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
