/*
Copyright © 2023 Gypsophlia <710852553@qq.com>
Generated by cobra-cli
*/
package cmd

import (
	"goto/interfaces"
	"goto/utils"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		sure := utils.InputBool("Are you sure to clear?")
		if sure {
			err := utils.WriteFile(dbFile, []interfaces.Task{})
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
