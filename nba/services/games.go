package nba

import (
	"encoding/json"
	"fmt"
	"nba-cli/entity"
	"net/http"
)

func GetGames(date string) (scbrd entity.Scoreboard, err error) {
	fmt.Println(date)
	GAME_URL := fmt.Sprintf("https://data.nba.net/prod/v1/%s/scoreboard.json", date)
	resp, err := http.Get(GAME_URL)

	if err != nil {
		panic(err)
		// log.Fatal("an error occurred, please try again")
	}

	defer resp.Body.Close()

	var scoreboard entity.Scoreboard
	if err := json.NewDecoder(resp.Body).Decode(&scoreboard); err != nil {
		panic(err)
		// log.Fatal("ooopsss! an error occurred, please try again")
	}
	return scoreboard, nil
}

func GetGameById()
