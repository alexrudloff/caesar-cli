package cmd

import (
	"fmt"

	"github.com/alexrudloff/caesar-cli/internal/client"
	"github.com/alexrudloff/caesar-cli/internal/output"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(brainstormCmd)
}

var brainstormCmd = &cobra.Command{
	Use:   "brainstorm [query]",
	Short: "Start a brainstorm session to get clarifying questions before research",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.New()
		if err != nil {
			output.Error("%v", err)
		}

		session, err := c.CreateBrainstorm(args[0])
		if err != nil {
			output.Error("%v", err)
		}

		fmt.Printf("Brainstorm session: %s\n\n", session.ID)
		for i, q := range session.Questions {
			fmt.Printf("%d. %s\n", i+1, q.Question)
			for k, v := range q.Options {
				fmt.Printf("   %s) %s\n", k, v)
			}
			fmt.Println()
		}
		fmt.Println("Use this session ID with: caesar research create --brainstorm", session.ID, "\"your query\"")
	},
}
