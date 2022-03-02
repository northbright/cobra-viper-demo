/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	showSmilingFace bool
	emoji           string
)

// hiCmd represents the hi command
var hiCmd = &cobra.Command{
	Use:   "hi [Name]",
	Short: "Say Hi",
	Long:  `Output "Hi, NAME"`,
	// hi command must have 1 arg: name to say hi.
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// Emoji flag is not set(--emoji), read emoji from config file.
		// default config file: $HOME/.demo.yaml(set in initConfig() in cmd/root.go)
		if len(emoji) == 0 {
			emoji = viper.GetString("emoji")
		}
		fmt.Printf("Hi, %v%v\n", name, emoji)
	},
}

func init() {
	rootCmd.AddCommand(hiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	hiCmd.Flags().StringVarP(&emoji, "emoji", "e", "", "Emoji to show")
}
