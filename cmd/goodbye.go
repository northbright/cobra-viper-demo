/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// goodbyeCmd represents the goodbye command
var goodbyeCmd = &cobra.Command{
	Use:   "goodbye [Name]",
	Short: "Say Goodbye",
	Long:  `Output "Goodbye, NAME"`,
	// goodbye command must have 1 arg: name to say goodbye.
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Goodbye, %v\n", name)
	},
}

func init() {
	rootCmd.AddCommand(goodbyeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goodbyeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goodbyeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
