/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package app_cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop web application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}

func init() {
	AppCmd.AddCommand(stopCmd)
}
