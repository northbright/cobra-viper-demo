/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	showMoon         bool
	showSleepingFace bool
)

// goodnightCmd represents the goodnight command
var goodnightCmd = &cobra.Command{
	Use:   "goodnight [Name]",
	Short: "Say Goodnight",
	Long:  `Ouput "Goodnight, Name"`,
	// goodnight command must have 1 arg: name to say goodnight.
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		s := "goodnight, %v"
		if showMoon {
			s += `🌙`
		}
		if showSleepingFace {
			s += `😴`
		}
		s += "\n"
		fmt.Printf(s, name)
	},
}

func init() {
	rootCmd.AddCommand(goodnightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goodnightCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goodnightCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	goodnightCmd.Flags().BoolVarP(&showMoon, "moon", "m", false, "show moon emoji")
	goodnightCmd.Flags().BoolVarP(&showSleepingFace, "sleepingface", "s", false, "show sleeping face emoji")
}
