package cmd

import (
	"os"

	app_cmd "github.com/binsabit/fasthttp-v1/cmd/app-cmd"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prosklad",
	Short: "This is Proskalad applicaiton",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func AddCmd() {
	rootCmd.AddCommand(app_cmd.AppCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	AddCmd()
}
