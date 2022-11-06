/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"

	"nba-cli/nba"

	"github.com/spf13/cobra"
)

var date = ""
var hasYesterday = false
var hasTomorrow = false

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Get the NBA schedule for a specific date",
	Run: func(cmd *cobra.Command, args []string) {
		date, _ = cmd.Flags().GetString("date")
		/* if hasYesterday {
			nba.FetchUpcomingGames(time.Now().AddDate(0, 0, -1), "y")
		}
		if hasTomorrow {
			nba.FetchUpcomingGames(time.Now().AddDate(0, 0, 1), "t")
		}
		// no date then get today's games
		if date == "" && !hasYesterday && !hasTomorrow {
			nba.FetchUpcomingGames(time.Now(), "")
		} else if date != "" {
			dateValue, _ := time.Parse("20060102", date)
			nba.FetchUpcomingGames(dateValue, "d")
		} */

	},
}

var gameIdCmd = &cobra.Command{
	Use:   "gameid",
	Short: "Get the NBA schedule for a specific date",
	Run: func(cmd *cobra.Command, args []string) {
		nba.GetGames(time.Now())
		/* for _, game := range scbrd.GameHeader {
			fmt.Println(game.GameID)
		} */
	},
}

func init() {
	rootCmd.AddCommand(gameCmd)
	rootCmd.AddCommand(gameIdCmd)
	rootCmd.PersistentFlags().StringVarP(&date, "date", "d", "", "Date to get the schedule for (YYYYMMDD)")
	rootCmd.PersistentFlags().BoolVarP(&hasYesterday, "yesterday", "y", false, "Get yesterday's games")
	rootCmd.PersistentFlags().BoolVarP(&hasTomorrow, "tomorrow", "t", false, "Get tomorrow's games")

	rootCmd.MarkFlagsMutuallyExclusive("yesterday", "tomorrow", "date")
}
