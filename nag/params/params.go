package params

import (
	"fmt"
	"strconv"
	"time"
)

type AheadBehind string

const (
	AheadOrBehind      AheadBehind = "Ahead or Behind"
	BehindOrTied                   = "Behind or Tied"
	AheadOrTied                    = "Ahead or Tied"
	DefaultAheadBehind             = AheadOrBehind
)

type ClutchTime string

const (
	Last5Minutes      ClutchTime = "Last 5 Minutes"
	Last4Minutes                 = "Last 4 Minutes"
	Last3Minutes                 = "Last 3 Minutes"
	Last2Minutes                 = "Last 2 Minutes"
	Last1Minute                  = "Last 1 Minute"
	Last30Seconds                = "Last 30 Seconds"
	Last10Seconds                = "Last 10 Seconds"
	DefaultClutchTime            = Last5Minutes
)

type Conference string

const (
	EastConference    Conference = "East"
	WestConference               = "West"
	NoneConference               = ""
	DefaultConference            = NoneConference
)

// ContextMeasureSimple(_ContextMeasure):
// ContextMeasureDetailed(_ContextMeasure):

type DefenseCategory string

const (
	Overall                DefenseCategory = "Overall"
	Threes                                 = "3 Pointers"
	Twos                                   = "2 Pointers"
	LessThan6ft                            = "Less Than 6Ft"
	LessThan10ft                           = "Less Than 10Ft"
	GreaterThan15ft                        = "Greater Than 15Ft"
	DefaultDefenseCategory                 = Overall
)

type Direction string

const (
	Asc  Direction = "ASC"
	Desc           = "DESC"

	DefaultDirection = Asc
)

type DistanceRange string

const (
	Range5ft             DistanceRange = "5ft Range"
	Range8ft                           = "8ft Range"
	ByZone                             = "By Zone"
	DefaultDistanceRange               = ByZone
)

type DivisionSimple string

const (
	Atlantic              DivisionSimple = "Atlantic"
	Central                              = "Central"
	Northwest                            = "Northwest"
	Pacific                              = "Pacific"
	Southeast                            = "Southeast"
	Southwest                            = "Southwest"
	DefaultDivisionSimple                = Atlantic
)

type Division string

const (
	EastDivision    Division = "East"
	WestDivision             = "West"
	DefaultDivision          = EastDivision
)

type GameScopeSimple string

const (
	Last10                 GameScopeSimple = "Last 10"
	Yesterday                              = "Yesterday"
	DefaultGameScopeSimple                 = Last10
)

// GameScopeDetailed(GameScopeSimple):
type GameScopeDetailed string

const (
	GameScopeDetailedSeason  GameScopeDetailed = "Season"
	GameScopeDetailedFinals                    = "Finals"
	DefaultGameScopeDetailed                   = GameScopeDetailedSeason
)

type GameSegment string

const (
	FirstHalf          GameSegment = "First Half"
	SecondHalf                     = "Second Half"
	Overtime                       = "Overtime"
	DefaultGameSegment             = FirstHalf
)

var (
	GroupQuantity        = func(i int) string { return strconv.Itoa(i) }
	DefaultGroupQuantity = "5"
)

var (
	LastNGames        = func(i int) string { return strconv.Itoa(i) }
	DefaultLastNGames = "0"
)

type leagueID struct{}

var LeagueID = leagueID{}

func (leagueID) NBA() string       { return "00" }
func (leagueID) ABA() string       { return "01" }
func (leagueID) WNBA() string      { return "10" }
func (leagueID) GLeague() string   { return "20" }
func (l leagueID) Default() string { return l.NBA() }

type Location string

const (
	Home            Location = "Home"
	Road                     = "Road"
	DefaultLocation          = Home
)

type MeasureTypeBase string

const (
	Base                   MeasureTypeBase = "Base"
	DefaultMeasureTypeBase                 = Base
)

// MeasureTypeSimple(MeasureTypeBase):
// MeasureTypePlayerGameLogs(MeasureTypeBase):

var (
	NumberOfGames        = func(i int) string { return strconv.Itoa(i) }
	DefaultNumberOfGames = "2147483647"
)

type Outcome string

const (
	Win            Outcome = "W"
	Loss                   = "L"
	DefaultOutcome         = Win
)

// PaceAdjust(_YesNo):
// PaceAdjustNo(_No):
// PlusMinus(_YesNo):
// PlusMinusNo(_No):

type period struct{}

var Period = period{}

func (period) All() string           { return "0" }
func (period) First() string         { return "1" }
func (period) Second() string        { return "2" }
func (period) Third() string         { return "3" }
func (period) Fourth() string        { return "4" }
func (period) Quarter(i int) string  { return strconv.Itoa(i) }
func (period) Overtime(i int) string { return strconv.Itoa(4 + i) }
func (p period) Default() string     { return p.All() }

type PerMode string

const (
	Totals         PerMode = "Totals"
	PerGame                = "PerGame"
	DefaultPerMode         = Totals
)

// PerMode36(PerModeSimple):
// PerMode48(PerModeSimple):
// PerModeTime(PerMode36, PerMode48):
// PerModeDetailed(PerModeTime):

type PlayerExperience string

const (
	Rookie                  PlayerExperience = "Rookie"
	Sophomore                                = "Sophomore"
	Veteran                                  = "Veteran"
	DefaultPlayerExperience                  = Rookie
)

type PlayerOrTeam string

const (
	Player              PlayerOrTeam = "Player"
	Team                             = "Team"
	DefaultPlayerOrTeam              = Team
)

type PlayerOrTeamAbbreviation string

const (
	P                               PlayerOrTeamAbbreviation = "P"
	T                                                        = "T"
	DefaultPlayerOrTeamAbbreviation                          = Team
)

type PlayerPosition string

const (
	Guard                 PlayerPosition = "Guard"
	Forward                              = "Forward"
	Center                               = "Center"
	DefaultPlayerPosition                = Guard
)

type PlayerPositionAbbreviation string

const (
	G                                 PlayerPositionAbbreviation = "G"
	F                                                            = "F"
	C                                                            = "C"
	GF                                                           = "G-F"
	FG                                                           = "F-G"
	FC                                                           = "F-C"
	CF                                                           = "C-F"
	DefaultPlayerPositionAbbreviation                            = G
)

type PlayerScope string

const (
	AllPlayers         PlayerScope = "All Players"
	Rookies                        = "Rookies"
	DefaultPlayerScope             = AllPlayers
)

// TodaysPlayers(_YesNo):
// ActivePlayers(_YesNo):

type PlayType string

const (
	Transition      PlayType = "Transition"
	Isolation                = "Isolation"
	PRBallHandler            = "PRBallHandler"
	PRRollMan                = "PRRollman"
	PostUp                   = "Postup"
	SpotUp                   = "Spotup"
	Handoff                  = "Handoff"
	Cut                      = "Cut"
	OffScreen                = "OffScreen"
	Putbacks                 = "OffRebound"
	Misc                     = "Misc"
	DefaultPlayType          = Transition
)

var (
	PointDiff        = func(i int) string { return strconv.Itoa(i) }
	DefaultPointDiff = "5"
)

type PtMeasureType string

const (
	SpeedDistance        PtMeasureType = "SpeedDistance"
	Rebounding                         = "Rebounding"
	Possessions                        = "Possessions"
	CatchShoot                         = "CatchShoot"
	PullUpShot                         = "PullUpShot"
	Defense                            = "Defense"
	Drives                             = "Drives"
	Passing                            = "Passing"
	ElbowTouch                         = "ElbowTouch"
	PostTouch                          = "PostTouch"
	PaintTouch                         = "PaintTouch"
	Efficiency                         = "Efficiency"
	DefaultPtMeasureType               = SpeedDistance
)

const (
	DefaultStartRange = "0"
	DefaultEndRange   = "0"
	DefaultRangeType  = "0"
)

// Rank(_YesNo):
// RankNo(_No):

type RunType string

const DefaultRunType RunType = "each second"

type Scope string

const (
	RSScope      Scope = "RS"
	SScope             = "S"
	RookiesScope       = "Rookies"
	DefaultScope       = SScope
)

// SeasonYear:
var (
	Season = func(t time.Time) string {
		cur := t.Year()
		if t.Month() <= 9 {
			cur = cur - 1
		}
		nxt := strconv.Itoa(cur + 1)[2:]
		return fmt.Sprintf("%d-%s", cur, nxt)
	}
	CurrentSeason = Season(time.Now())
)

// SeasonAll(Season):
// SeasonAll_Time(Season):
// SeasonAllTime(Season):
// SeasonID(SeasonYear):

type SeasonType string

const (
	Regular           SeasonType = "Regular Season"
	PreSeason                    = "Pre Season"
	Playoffs                     = "Playoff"
	AllStar                      = "All Star"
	DefaultSeasonType            = Regular
)

type SeasonSegment string

const (
	PostAllStar          SeasonSegment = "Post All-Star"
	PreAllStar                         = "Pre All-Star"
	DefaultSeasonSegment               = PostAllStar
)

type ShotClockRange string

const (
	Range2224             ShotClockRange = "24-22"
	Range1822                            = "22-18 Very Early"
	Range1518                            = "18-15 Early"
	Range715                             = "15-7 Average"
	Range47                              = "7-4 Late"
	Range04                              = "4-0 Very Late"
	ShotClockOff                         = "ShotClock Off"
	Empty                                = ""
	DefaultShotClockRange                = ShotClockOff
)

/*
func CalculateRange(i int64) ShotClockRange {
	switch {
	case i > 24, i <= 0:
		return Empty
	case 22 < i <= 24:
		return Range2224
	case 18 < i <= 22:
		return Range1822
	case 15 < i <= 18:
		return Range1518
	case 7 < i <= 15:
		return Range715
	case 4 < i <= 7:
		return Range47
	case 0 < i <= 4:
		return Range04
	}
}
*/

type Sorter string

const (
	SorterFGM     Sorter = "FGM"
	SorterFGA            = "FGA"
	SorterFG_PCT         = "FG_PCT"
	SorterFG3M           = "FG3M"
	SorterFG3A           = "FG3A"
	SorterFG3_PCT        = "FG3_PCT"
	SorterFTM            = "FTM"
	SorterFTA            = "FTA"
	SorterFT_PCT         = "FT_PCT"
	SorterOREB           = "OREB"
	SorterDREB           = "DREB"
	SorterAST            = "AST"
	SorterSTL            = "STL"
	SorterBLK            = "BLK"
	SorterTOV            = "TOV"
	SorterREB            = "REB"
	SorterPTS            = "PTS"
	SorterDate           = "DATE"
	DefaultSorter        = SorterDate
)

type StarterBench string

const (
	Starters            StarterBench = "Starters"
	Bench                            = "Bench"
	DefaultStarterBench              = Starters
)

type Stat string

const (
	Points           Stat = "PTS"
	Rebounds              = "REB"
	Assists               = "AST"
	FieldGoalPercent      = "FG_PCT"
	FreeThrowPercent      = "FT_PCT"
	ThreesPercent         = "FG3_PCT"
	Steals                = "STL"
	Blocks                = "BLK"
	DefaultStat           = Points
)

type StatCategory string

const (
	StatCategoryPoints           StatCategory = "Points"
	StatCategoryRebounds                      = "Rebounds"
	StatCategoryAssists                       = "Assists"
	StatCategoryDefense                       = "Defense"
	StatCategoryClutch                        = "Clutch"
	StatCategoryPlaymaking                    = "Playmaking"
	StatCategoryEfficiency                    = "Efficiency"
	StatCategoryFastBreak                     = "Fast Break"
	StatCategoryScoringBreakdown              = "Scoring Breakdown"
	DefaultStatCategory                       = StatCategoryPoints
)

type StatCategoryAbbreviation string

const (
	PTS                             StatCategoryAbbreviation = "PTS"
	FGM                                                      = "FGM"
	FGA                                                      = "FGA"
	FG_PCT                                                   = "FG_PCT"
	FG3M                                                     = "FG3M"
	FG3A                                                     = "FG3A"
	FG3_PCT                                                  = "FG3_PCT"
	FTM                                                      = "FTM"
	FTA                                                      = "FTA"
	OREB                                                     = "OREB"
	DREB                                                     = "DREB"
	AST                                                      = "AST"
	STL                                                      = "STL"
	BLK                                                      = "BLK"
	TOV                                                      = "TOV"
	REB                                                      = "REB"
	DefaultStatCategoryAbbreviation                          = PTS
)

type StatType string

const (
	Traditional     StatType = "Traditional"
	Advanced                 = "Advanced"
	Tracking                 = "Tracking"
	DefaultStatType          = Traditional
)

type TypeGrouping string

const (
	Offensive           TypeGrouping = "offensive"
	Defensive                        = "defensive"
	DefaultTypeGrouping              = Offensive
)
