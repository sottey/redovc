package cmd

import (
	"strings"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		setStatusCmdDesc    = "Sets the status of a todo item"
		setStatusCmdExample = `  To add a "blocked" status to a todo:
    ultodo status 33 blocked
    ultodo s 33 blocked

  You can remove a status by setting a status to "none".  Example:
    ultodo s 33 none`

		setStatusCmdLongDesc = `Sets the status of a todo item.
  A status should be a single lower-case word, e.g. "now", "blocked", or "waiting".

  For more info, see https://ultodo.io/docs/cli/managing_tasks/#handling-todo-statuses`
	)

	var setStatusCmd = &cobra.Command{
		Use:     "status [id] <status>",
		Aliases: []string{"s"},
		Example: setStatusCmdExample,
		Long:    setStatusCmdLongDesc,
		Short:   setStatusCmdDesc,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().SetTodoStatus(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(setStatusCmd)
}
