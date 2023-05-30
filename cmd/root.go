/*
Copyright © 2023 Gypsophlia <710852553@qq.com>
Generated by cobra-cli
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// global variant
var dbFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goto",
	Short: "GOTO: A tool to manage command tasks.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	dir, _ := os.UserHomeDir()
	dbFile = dir + "./.goto.db"
	// rootCmd.PersistentFlags().StringVar(&dbFile, "db", dir+"/.goto.db", "database file (default is .goto.db)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
