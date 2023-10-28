package cmd

import (
	"strings"

	redovc "github.com/sottey/redo.vc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		archiveCompletedTodo bool
		long                 = `Completes or un-completes a todo.`
		completeCmdExample   = `
  Complete a todo with id 33:
    redovc complete 33
    redovc c 33

  Complete a todo with id 33 and archive it:
    redovc uncomplete 33 --archive

  Uncompletes todo with id 33.
    redovc uncomplete 33
    redovc uc 33`
	)

	var completeCmd = &cobra.Command{
		Use:     "complete [id]",
		Aliases: []string{"c"},
		Example: completeCmdExample,
		Short:   "Completes a todo.",
		Long:    long,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().CompleteTodo(strings.Join(args, " "), archiveCompletedTodo)
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
			redovc.NewApp().UncompleteTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().BoolVarP(&archiveCompletedTodo, "archive", "", false, "Archives a completed todo automatically")
	rootCmd.AddCommand(uncompleteCmd)
}
