package cmd

import (
	"strings"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		archiveCompletedTodo bool
		long                 = `Completes or un-completes a todo.

For more info, see https://ultodo.io/docs/cli/managing_tasks`
		completeCmdExample = `
  Complete a todo with id 33:
    ultodo complete 33
    ultodo c 33

  Complete a todo with id 33 and archive it:
    ultodo uncomplete 33 --archive

  Uncompletes todo with id 33.
    ultodo uncomplete 33
    ultodo uc 33`
	)

	var completeCmd = &cobra.Command{
		Use:     "complete [id]",
		Aliases: []string{"c"},
		Example: completeCmdExample,
		Short:   "Completes a todo.",
		Long:    long,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().CompleteTodo(strings.Join(args, " "), archiveCompletedTodo)
		},
	}

	var uncompleteCmd = &cobra.Command{
		Use:     "uncomplete [id]",
		Aliases: []string{"uc"},
		Example: completeCmdExample,
		Short:   "Un-completes a todo.",
		Long:    long,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().UncompleteTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().BoolVarP(&archiveCompletedTodo, "archive", "", false, "Archives a completed todo automatically")
	rootCmd.AddCommand(uncompleteCmd)
}
