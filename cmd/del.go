/*
Copyright © 2023 Gypsophlia <710852553@qq.com>
Generated by cobra-cli
*/
package cmd

import (
	"fmt"
	"goto/utils"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "del a task by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := utils.ReadFile(dbFile)
		if err != nil {
			return err
		}
		index := utils.FindIndex(args[0], tasks)
		if index == -1 {
			return fmt.Errorf("the name of task doesn't exist")
		}
		tasks = append(tasks[:index], tasks[index+1:]...)
		err = utils.WriteFile(dbFile, tasks)
		return err
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}