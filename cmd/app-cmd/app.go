package app_cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// appCmd represents the app command
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "commads ralated to running and stopping app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("app called")
	},
}

func init() {

}
