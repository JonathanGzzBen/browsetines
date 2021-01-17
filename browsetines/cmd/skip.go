package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(skipCmd)
}

var skipCmd = &cobra.Command{
	Use:   "skip <id> [<n>]",
	Short: "Skip activity",
	Long:  `Skip "n" activities in routine with matching ID.`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Skipping activities...")
	},
}
