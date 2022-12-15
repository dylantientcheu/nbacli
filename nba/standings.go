package nba

import (
	"github.com/mitchellh/mapstructure"

	"github.com/dylantientcheu/nbacli/nag"
)

type Standing struct {
	LeagueID                string
	SeasonID                string
	TeamID                  int64
	TeamCity                string
	TeamName                string
	TeamSlug                string
	Conference              nag.Conference
	ConferenceRecord        string
	PlayoffRank             int64
	ClinchIndicator         string
	Division                string
	DivisionRecord          string
	DivisionRank            int64
	WINS                    int64
	Losses                  int64
	WinPCT                  float64
	LeagueRank              int64
	Record                  string
	Home                    string
	Road                    string
	L10                     string
	Last10Home              string
	Last10Road              string
	Ot                      string
	ThreePTSOrLess          string
	TenPTSOrMore            string
	LongHomeStreak          int64
	StrLongHomeStreak       string
	LongRoadStreak          int64
	StrLongRoadStreak       string
	LongWinStreak           int64
	LongLossStreak          int64
	CurrentHomeStreak       int64
	StrCurrentHomeStreak    string
	CurrentRoadStreak       int64
	StrCurrentRoadStreak    string
	CurrentStreak           int64
	StrCurrentStreak        string
	ConferenceGamesBack     float64
	DivisionGamesBack       float64
	ClinchedConferenceTitle int64
	ClinchedDivisionTitle   int64
	ClinchedPlayoffBirth    int64
	ClinchedPlayIn          int64
	EliminatedConference    int64
	EliminatedDivision      int64
	AheadAtHalf             string
	BehindAtHalf            string
	TiedAtHalf              string
	AheadAtThird            string
	BehindAtThird           string
	TiedAtThird             string
	Score100PTS             string
	OppScore100PTS          string
	OppOver500              string
	LeadInFGPCT             string
	LeadInReb               string
	FewerTurnovers          string
	PointsPG                float64
	OppPointsPG             float64
	DiffPointsPG            float64
	TotalPoints             int64
	OppTotalPoints          int64
	DiffTotalPoints         int64
}

type StandingsRepository struct {
}

func (g *StandingsRepository) GetSeasonStandings() ([]Standing, []Standing) {
	sbv2 := nag.NewLeagueStandingsV3()
	err := sbv2.Get()
	if err != nil {
		panic(err)
	}
	if sbv2.Response == nil {
		panic("no response")
	}

	n := nag.Map(*sbv2.Response)
	var result nag.LeagueStandingsResponse
	mapstructure.Decode(n, &result)

	easternConference := make([]Standing, 0, len(result.Standings)/2)
	westernConference := make([]Standing, 0, len(result.Standings)/2)

	for _, v := range result.Standings {
		var standing Standing
		standing.TeamID = v.TeamID
		standing.PlayoffRank = v.PlayoffRank
		standing.TeamName = v.TeamName
		standing.WINS = v.WINS
		standing.Losses = v.Losses
		standing.WinPCT = v.WinPCT
		standing.Home = v.Home
		standing.Road = v.Road
		standing.Conference = v.Conference
		standing.ConferenceRecord = v.ConferenceRecord
		standing.PointsPG = v.PointsPG
		standing.OppPointsPG = v.OppPointsPG
		standing.DiffPointsPG = v.DiffPointsPG
		standing.StrCurrentStreak = v.StrCurrentStreak
		standing.L10 = v.L10

		if v.Conference == nag.East {
			easternConference = append(easternConference, standing)
		} else {
			westernConference = append(westernConference, standing)
		}
	}

	return easternConference, westernConference
}
