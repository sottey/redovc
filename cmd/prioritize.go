package cmd

import (
	"strings"

	"github.com/sottey/redo.vc/redovc"
	"github.com/spf13/cobra"
)

func init() {
	var (
		example = `  To prioritize a todo with id 33:
    redovc prioritize 33
    redovc p 33

  To un-prioritize a todo with an id 33:
    redovc unprioritize 33
    redovc up 33`

		long = `Prioritize and un-prioritize todos.

  Todos with the priority flag will be highlighted, and will be at the top of your list.`
	)

	var prioritizeCmd = &cobra.Command{
		Use:     "prioritize [id]",
		Aliases: []string{"p"},
		Example: example,
		Long:    long,
		Short:   "Prioritize a todo.",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().PrioritizeTodo(strings.Join(args, " "))
		},
	}

	var unprioritizeCmd = &cobra.Command{
		Use:     "unprioritize [id]",
		Aliases: []string{"up"},
		Example: example,
		Long:    long,
		Short:   "Un-prioritize a todo.",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().UnprioritizeTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(prioritizeCmd)
	rootCmd.AddCommand(unprioritizeCmd)
}
