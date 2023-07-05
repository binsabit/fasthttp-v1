/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package app_cmd

import (
	"github.com/binsabit/fasthttp-v1/app"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "start",
	Short: "start application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app.StartApp()
	},
}

func init() {
	AppCmd.AddCommand(runCmd)
}
