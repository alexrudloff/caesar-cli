package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caesar",
	Short: "CLI for the Caesar research API",
	Long:  "A command-line interface for caesar.org's research API. Set CAESAR_API_KEY to authenticate.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
