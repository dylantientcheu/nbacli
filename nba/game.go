package nba

import (
	"github.com/mitchellh/mapstructure"

	"nbacli/nag"
)

type GameStat struct {
	GameID           string
	TeamID           int64
	TeamAbbreviation string
	TeamCity         string
	PlayerID         int64
	PlayerName       string
	Nickname         string
	StartPosition    string
	Comment          string
	Min              string
	Fgm              int64
	Fga              int64
	FgPct            float64
	Fg3M             int64
	Fg3A             int64
	Fg3Pct           float64
	Ftm              int64
	Fta              int64
	FtPct            float64
	Oreb             int64
	Dreb             int64
	Reb              int64
	AST              int64
	Stl              int64
	Blk              int64
	To               int64
	Pf               int64
	Pts              int64
	PlusMinus        int64
	TeamName         string
	StartersBench    string
}

type BoxScoreRepository struct {
}

func (g *BoxScoreRepository) GetSingleGameStats(gameID string) []GameStat {
	sbv2 := nag.NewBoxScoreTraditionalV2(gameID)
	err := sbv2.Get()
	if err != nil {
		panic(err)
	}
	if sbv2.Response == nil {
		panic("no response")
	}

	n := nag.Map(*sbv2.Response)
	var result nag.BoxScoreTraditionalResponse
	mapstructure.Decode(n, &result)

	stats := make([]GameStat, 0, len(result.PlayerStats))

	for _, v := range result.PlayerStats {
		var gameStat GameStat
		gameStat.GameID = v.GameID
		gameStat.TeamID = v.TeamID
		gameStat.TeamAbbreviation = v.TeamAbbreviation
		gameStat.TeamCity = v.TeamCity
		gameStat.PlayerID = v.PlayerID
		gameStat.PlayerName = v.PlayerName
		gameStat.Nickname = v.Nickname
		gameStat.StartPosition = v.StartPosition
		gameStat.Comment = v.Comment
		gameStat.Min = v.Min
		gameStat.Fgm = v.Fgm
		gameStat.Fga = v.Fga
		gameStat.FgPct = v.FgPct
		gameStat.Fg3M = v.Fg3M
		gameStat.Fg3A = v.Fg3A
		gameStat.Fg3Pct = v.Fg3Pct
		gameStat.Ftm = v.Ftm
		gameStat.Fta = v.Fta
		gameStat.FtPct = v.FtPct
		gameStat.Oreb = v.Oreb
		gameStat.Dreb = v.Dreb
		gameStat.Reb = v.Reb
		gameStat.AST = v.AST
		gameStat.Stl = v.Stl
		gameStat.Blk = v.Blk
		gameStat.To = v.To
		gameStat.Pf = v.Pf
		gameStat.Pts = v.Pts
		gameStat.PlusMinus = v.PlusMinus
		gameStat.TeamName = v.TeamName
		gameStat.StartersBench = v.StartersBench

		stats = append(stats, gameStat)
	}

	return stats
}
