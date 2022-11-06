package nba

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"

	"nba-cli/nag"
)

func GetGames(date time.Time) (scbrd nag.Response) {
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
	fmt.Printf("%#v", result.GameHeader)

	for _, v := range result.GameHeader {
		fmt.Println(v.GameID)
	}
	return *sbv2.Response
}

func GetGameById(gameId string) {
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
