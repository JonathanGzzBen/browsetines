package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var rootCmd = &cobra.Command{
	Use:   "browsetines",
	Short: "browsetines lets you program routines",
	Long: `Browsetines is a command-line tool to program and control
how much time you invest in certain activities.`,
}

// Execute runs the root command
func Execute() {
	err := doc.GenMarkdownTree(rootCmd, "docs")
	if err != nil {
		log.Fatal(err)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
