package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:     "remove <id>",
	Short:   "Remove routine",
	Long:    "Remove routine with matching ID.",
	Aliases: []string{"rm", "delete", "del"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing routine...")
	},
}
