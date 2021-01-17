package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var prompt bool
var delay time.Duration
var urls []string

func init() {
	editAddCmd.Flags().BoolVarP(
		&prompt,
		"prompt",
		"p",
		false,
		"prompt before starting activity")
	editAddCmd.Flags().StringSliceVarP(
		&urls,
		"urls",
		"u",
		[]string{},
		"urls to open when activity starts")

	editAddCmd.Flags().DurationVarP(
		&delay,
		"delay",
		"d",
		time.Duration(0),
		"delay before starting activity")

	editCmd.AddCommand(editAddCmd)
	editCmd.AddCommand(editRemoveCmd)
	rootCmd.AddCommand(editCmd)

}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit routine",
	Long:  `Edit a routine by adding or removing activities.`,
}

var editAddCmd = &cobra.Command{
	Use:   "add <id> <name> [<position>] [<time>]",
	Short: "Add an activity",
	Long: `Add an activity with provided name to routine
with given ID.
If position is not specified or is -1, activity will be
inserted at last position.
If you provide a time, a stopwatch will run next
activity in routine when it reaches zero.
Flag prompt overrides delay.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adding a new activity...")
	},
}

var editRemoveCmd = &cobra.Command{
	Use:   "remove <id> <position>",
	Short: "Remove activity",
	Long: `Remove activity in specified position
from routine with matching ID.`,
	Aliases: []string{"rm", "delete", "del"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing activity...")
	},
}
