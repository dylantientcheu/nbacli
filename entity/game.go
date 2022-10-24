package entity

import "time"

type Today struct {
	Internal struct {
		PubDateTime             string `json:"pubDateTime"`
		IgorPath                string `json:"igorPath"`
		Xslt                    string `json:"xslt"`
		XsltForceRecompile      string `json:"xsltForceRecompile"`
		XsltInCache             string `json:"xsltInCache"`
		XsltCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
		XsltTransformTimeMillis string `json:"xsltTransformTimeMillis"`
		ConsolidatedDomKey      string `json:"consolidatedDomKey"`
		EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
	} `json:"_internal"`
	TeamSitesOnly struct {
		SeasonStage    int    `json:"seasonStage"`
		SeasonYear     int    `json:"seasonYear"`
		RosterYear     int    `json:"rosterYear"`
		StatsStage     int    `json:"statsStage"`
		StatsYear      int    `json:"statsYear"`
		DisplayYear    string `json:"displayYear"`
		LastPlayByPlay string `json:"lastPlayByPlay"`
		AllPlayByPlay  string `json:"allPlayByPlay"`
		PlayerMatchup  string `json:"playerMatchup"`
		Series         string `json:"series"`
	} `json:"teamSitesOnly"`
	SeasonScheduleYear int  `json:"seasonScheduleYear"`
	ShowPlayoffsClinch bool `json:"showPlayoffsClinch"`
	Links              struct {
		AnchorDate                  string `json:"anchorDate"`
		CurrentDate                 string `json:"currentDate"`
		Calendar                    string `json:"calendar"`
		TodayScoreboard             string `json:"todayScoreboard"`
		CurrentScoreboard           string `json:"currentScoreboard"`
		Teams                       string `json:"teams"`
		Scoreboard                  string `json:"scoreboard"`
		LeagueRosterPlayers         string `json:"leagueRosterPlayers"`
		AllstarRoster               string `json:"allstarRoster"`
		LeagueRosterCoaches         string `json:"leagueRosterCoaches"`
		LeagueSchedule              string `json:"leagueSchedule"`
		LeagueConfStandings         string `json:"leagueConfStandings"`
		LeagueDivStandings          string `json:"leagueDivStandings"`
		LeagueUngroupedStandings    string `json:"leagueUngroupedStandings"`
		LeagueMiniStandings         string `json:"leagueMiniStandings"`
		LeagueTeamStatsLeaders      string `json:"leagueTeamStatsLeaders"`
		LeagueLastFiveGameTeamStats string `json:"leagueLastFiveGameTeamStats"`
		PreviewArticle              string `json:"previewArticle"`
		RecapArticle                string `json:"recapArticle"`
		GameBookPdf                 string `json:"gameBookPdf"`
		Boxscore                    string `json:"boxscore"`
		MiniBoxscore                string `json:"miniBoxscore"`
		Pbp                         string `json:"pbp"`
		LeadTracker                 string `json:"leadTracker"`
		PlayerGameLog               string `json:"playerGameLog"`
		PlayerProfile               string `json:"playerProfile"`
		PlayerUberStats             string `json:"playerUberStats"`
		TeamSchedule                string `json:"teamSchedule"`
		TeamsConfig                 string `json:"teamsConfig"`
		TeamRoster                  string `json:"teamRoster"`
		TeamsConfigYear             string `json:"teamsConfigYear"`
		TeamScheduleYear            string `json:"teamScheduleYear"`
		TeamLeaders                 string `json:"teamLeaders"`
		TeamScheduleYear2           string `json:"teamScheduleYear2"`
		TeamLeaders2                string `json:"teamLeaders2"`
		TeamICS                     string `json:"teamICS"`
		TeamICS2                    string `json:"teamICS2"`
		PlayoffsBracket             string `json:"playoffsBracket"`
		PlayoffSeriesLeaders        string `json:"playoffSeriesLeaders"`
		UniversalLinkMapping        string `json:"universalLinkMapping"`
		TicketLink                  string `json:"ticketLink"`
	} `json:"links"`
}

type Game struct {
	SeasonStageID int
	SeasonYear    string
	LeagueName    string
	GameID        string
	Arena         struct {
		Name       string
		IsDomestic bool
		City       string
		StateAbbr  string
		Country    string
	}
	IsGameActivated       bool
	StatusNum             int
	ExtendedStatusNum     int
	StartTimeEastern      string
	StartTimeUTC          time.Time
	EndTimeUTC            time.Time
	StartDateEastern      string
	HomeStartDate         string
	HomeStartTime         string
	VisitorStartDate      string
	VisitorStartTime      string
	GameURLCode           string
	Clock                 string
	IsBuzzerBeater        bool
	IsPreviewArticleAvail bool
	IsRecapArticleAvail   bool
	Nugget                struct {
		Text string
	}
	Attendance string
	Tickets    struct {
		MobileApp    string
		DesktopWeb   string
		MobileWeb    string
		LeagGameInfo string
		LeagTix      string
	}
	HasGameBookPdf bool
	IsStartTimeTBD bool
	IsNeutralVenue bool
	GameDuration   struct {
		Hours   string
		Minutes string
	}
	Period struct {
		Current       int
		Type          int
		MaxRegular    int
		IsHalftime    bool
		IsEndOfPeriod bool
	}
	VTeam struct {
		TeamID     string
		TriCode    string
		Win        string
		Loss       string
		SeriesWin  string
		SeriesLoss string
		Score      string
		Linescore  []struct {
			Score string
		}
	}
	HTeam struct {
		TeamID     string
		TriCode    string
		Win        string
		Loss       string
		SeriesWin  string
		SeriesLoss string
		Score      string
		Linescore  []struct {
			Score string
		}
	}
	Watch struct {
		Broadcast struct {
			Broadcasters struct {
				National []interface{}
				Canadian []struct {
					ShortName string
					LongName  string
				}
				VTeam []struct {
					ShortName string
					LongName  string
				}
				HTeam []struct {
					ShortName string
					LongName  string
				}
				SpanishHTeam    []interface{}
				SpanishVTeam    []interface{}
				SpanishNational []interface{}
			}
			Video struct {
				RegionalBlackoutCodes string
				CanPurchase           bool
				IsLeaguePass          bool
				IsNationalBlackout    bool
				IsTNTOT               bool
				IsVR                  bool
				TntotIsOnAir          bool
				IsNextVR              bool
				IsNBAOnTNTVR          bool
				IsMagicLeap           bool
				IsOculusVenues        bool
				Streams               []struct {
					StreamType            string
					IsOnAir               bool
					DoesArchiveExist      bool
					IsArchiveAvailToWatch bool
					StreamID              string
					Duration              int
				}
				DeepLink []interface{}
			}
			Audio struct {
				National struct {
					Streams []struct {
						Language string
						IsOnAir  bool
						StreamID string
					}
					Broadcasters []interface{}
				}
				VTeam struct {
					Streams []struct {
						Language string
						IsOnAir  bool
						StreamID string
					}
					Broadcasters []struct {
						ShortName string
						LongName  string
					}
				}
				HTeam struct {
					Streams []struct {
						Language string
						IsOnAir  bool
						StreamID string
					}
					Broadcasters []struct {
						ShortName string
						LongName  string
					}
				}
			}
		}
	}
}

type Scoreboard struct {
	Internal struct {
		PubDateTime             string `json:"pubDateTime"`
		IgorPath                string `json:"igorPath"`
		RouteName               string `json:"routeName"`
		RouteValue              string `json:"routeValue"`
		Xslt                    string `json:"xslt"`
		XsltForceRecompile      string `json:"xsltForceRecompile"`
		XsltInCache             string `json:"xsltInCache"`
		XsltCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
		XsltTransformTimeMillis string `json:"xsltTransformTimeMillis"`
		ConsolidatedDomKey      string `json:"consolidatedDomKey"`
		EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
	} `json:"_internal"`
	NumGames int `json:"numGames"`
	Games    []Game
}
