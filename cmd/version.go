package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the macshift",
	Long:  `This command prints the current version of the macshift.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.1.1")
	},
}
