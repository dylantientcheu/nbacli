package nag

type NewBoxScoreAdvancedResponse struct {
	PlayerStats []Stat `mapstructure:"PlayerStats"`
	TeamStats   []Stat `mapstructure:"TeamStats"`
}

type Stat struct {
	GameID           string  `mapstructure:"GAME_ID"`
	TeamID           int64   `mapstructure:"TEAM_ID"`
	TeamAbbreviation string  `mapstructure:"TEAM_ABBREVIATION"`
	TeamCity         string  `mapstructure:"TEAM_CITY"`
	PlayerID         int64   `mapstructure:"PLAYER_ID,omitempty"`
	PlayerName       string  `mapstructure:"PLAYER_NAME,omitempty"`
	Nickname         string  `mapstructure:"NICKNAME,omitempty"`
	StartPosition    string  `mapstructure:"START_POSITION,omitempty"`
	Comment          string  `mapstructure:"COMMENT,omitempty"`
	Min              string  `mapstructure:"MIN"`
	EOffRating       float64 `mapstructure:"E_OFF_RATING"`
	OffRating        float64 `mapstructure:"OFF_RATING"`
	EDefRating       float64 `mapstructure:"E_DEF_RATING"`
	DefRating        float64 `mapstructure:"DEF_RATING"`
	ENetRating       float64 `mapstructure:"E_NET_RATING"`
	NetRating        float64 `mapstructure:"NET_RATING"`
	ASTPct           float64 `mapstructure:"AST_PCT"`
	ASTTov           float64 `mapstructure:"AST_TOV"`
	ASTRatio         float64 `mapstructure:"AST_RATIO"`
	OrebPct          float64 `mapstructure:"OREB_PCT"`
	DrebPct          float64 `mapstructure:"DREB_PCT"`
	RebPct           float64 `mapstructure:"REB_PCT"`
	TmTovPct         float64 `mapstructure:"TM_TOV_PCT"`
	EfgPct           float64 `mapstructure:"EFG_PCT"`
	TsPct            float64 `mapstructure:"TS_PCT"`
	UsgPct           float64 `mapstructure:"USG_PCT"`
	EUsgPct          float64 `mapstructure:"E_USG_PCT"`
	EPace            float64 `mapstructure:"E_PACE"`
	Pace             float64 `mapstructure:"PACE"`
	PacePer40        float64 `mapstructure:"PACE_PER40"`
	Poss             int64   `mapstructure:"POSS"`
	Pie              float64 `mapstructure:"PIE"`
	TeamName         string  `mapstructure:"TEAM_NAME,omitempty"`
	ETmTovPct        float64 `mapstructure:"E_TM_TOV_PCT,omitempty"`
}

type BoxScoreSummaryResponse struct {
	GameSummary     []GameSummary    `mapstructure:"GameSummary"`
	OtherStats      []OtherStat      `mapstructure:"OtherStats"`
	Officials       []Official       `mapstructure:"Officials"`
	InactivePlayers []InactivePlayer `mapstructure:"InactivePlayers"`
	GameInfo        []GameInfo       `mapstructure:"GameInfo"`
	LineScore       []LineScore      `mapstructure:"LineScore"`
	LastMeeting     []LastMeeting    `mapstructure:"LastMeeting"`
	SeasonSeries    []SeasonSery     `mapstructure:"SeasonSeries"`
	AvailableVideo  []AvailableVideo `mapstructure:"AvailableVideo"`
}

type AvailableVideo struct {
	GameID             string `mapstructure:"GAME_ID"`
	VideoAvailableFlag int64  `mapstructure:"VIDEO_AVAILABLE_FLAG"`
	PtAvailable        int64  `mapstructure:"PT_AVAILABLE"`
	PtXyzAvailable     int64  `mapstructure:"PT_XYZ_AVAILABLE"`
	WhStatus           int64  `mapstructure:"WH_STATUS"`
	HustleStatus       int64  `mapstructure:"HUSTLE_STATUS"`
	HistoricalStatus   int64  `mapstructure:"HISTORICAL_STATUS"`
}

type GameInfo struct {
	GameDate   string `mapstructure:"GAME_DATE"`
	Attendance int64  `mapstructure:"ATTENDANCE"`
	GameTime   string `mapstructure:"GAME_TIME"`
}

type GameSummary struct {
	GameDateEst                   string      `mapstructure:"GAME_DATE_EST"`
	GameSequence                  int64       `mapstructure:"GAME_SEQUENCE"`
	GameID                        string      `mapstructure:"GAME_ID"`
	GameStatusID                  int64       `mapstructure:"GAME_STATUS_ID"`
	GameStatusText                string      `mapstructure:"GAME_STATUS_TEXT"`
	Gamecode                      string      `mapstructure:"GAMECODE"`
	HomeTeamID                    int64       `mapstructure:"HOME_TEAM_ID"`
	VisitorTeamID                 int64       `mapstructure:"VISITOR_TEAM_ID"`
	Season                        string      `mapstructure:"SEASON"`
	LivePeriod                    int64       `mapstructure:"LIVE_PERIOD"`
	LivePCTime                    string      `mapstructure:"LIVE_PC_TIME"`
	NatlTvBroadcasterAbbreviation interface{} `mapstructure:"NATL_TV_BROADCASTER_ABBREVIATION"`
	LivePeriodTimeBcast           string      `mapstructure:"LIVE_PERIOD_TIME_BCAST"`
	WhStatus                      int64       `mapstructure:"WH_STATUS"`
}

type InactivePlayer struct {
	PlayerID         int64  `mapstructure:"PLAYER_ID"`
	FirstName        string `mapstructure:"FIRST_NAME"`
	LastName         string `mapstructure:"LAST_NAME"`
	JerseyNum        string `mapstructure:"JERSEY_NUM"`
	TeamID           int64  `mapstructure:"TEAM_ID"`
	TeamCity         string `mapstructure:"TEAM_CITY"`
	TeamName         string `mapstructure:"TEAM_NAME"`
	TeamAbbreviation string `mapstructure:"TEAM_ABBREVIATION"`
}

type Official struct {
	OfficialID int64  `mapstructure:"OFFICIAL_ID"`
	FirstName  string `mapstructure:"FIRST_NAME"`
	LastName   string `mapstructure:"LAST_NAME"`
	JerseyNum  string `mapstructure:"JERSEY_NUM"`
}

type OtherStat struct {
	LeagueID         string `mapstructure:"LEAGUE_ID"`
	TeamID           int64  `mapstructure:"TEAM_ID"`
	TeamAbbreviation string `mapstructure:"TEAM_ABBREVIATION"`
	TeamCity         string `mapstructure:"TEAM_CITY"`
	PtsPaint         int64  `mapstructure:"PTS_PAINT"`
	Pts2NdChance     int64  `mapstructure:"PTS_2ND_CHANCE"`
	PtsFb            int64  `mapstructure:"PTS_FB"`
	LargestLead      int64  `mapstructure:"LARGEST_LEAD"`
	LeadChanges      int64  `mapstructure:"LEAD_CHANGES"`
	TimesTied        int64  `mapstructure:"TIMES_TIED"`
	TeamTurnovers    int64  `mapstructure:"TEAM_TURNOVERS"`
	TotalTurnovers   int64  `mapstructure:"TOTAL_TURNOVERS"`
	TeamRebounds     int64  `mapstructure:"TEAM_REBOUNDS"`
	PtsOffTo         int64  `mapstructure:"PTS_OFF_TO"`
}

type SeasonSery struct {
	GameID string `mapstructure:"GAME_ID"`
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
	Pts              int         `mapstructure:"PTS"`
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
