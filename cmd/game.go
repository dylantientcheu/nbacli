/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"

	"github.com/blurdylan/go-nba/nba"
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
		if hasYesterday {
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
		}

	},
}

func init() {
	rootCmd.AddCommand(gameCmd)
	rootCmd.PersistentFlags().StringVarP(&date, "date", "d", "", "Date to get the schedule for (YYYYMMDD)")
	rootCmd.PersistentFlags().BoolVarP(&hasYesterday, "yesterday", "y", false, "Get yesterday's games")
	rootCmd.PersistentFlags().BoolVarP(&hasTomorrow, "tomorrow", "t", false, "Get tomorrow's games")

	rootCmd.MarkFlagsMutuallyExclusive("yesterday", "tomorrow", "date")
	// cannot call both flags at the same time

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
