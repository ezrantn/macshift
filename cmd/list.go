package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ezrantn/macshift/adapter"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all network adapters",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Writer = os.Stdout

		s.Start()

		adapters, err := adapter.ListAdapters()
		if err != nil {
			s.Stop()
			fmt.Fprintf(os.Stderr, "Error: Unable to list network adapters. %v\n", err)
			fmt.Fprintln(os.Stderr, "Please check if you have the necessary permissions or try running as administrator.")
			os.Exit(1)
		}

		s.Stop()
		fmt.Print("\r")

		fmt.Println("Available network adapters:")
		for _, adapter := range adapters {
			fmt.Printf("\nName: %s\nDescription: %s\nMAC: %s\n",
				adapter.Name, adapter.Description, adapter.MacAddress)
		}
	},
}
