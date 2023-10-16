package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		editCmdDesc = "Edits todos"
		longDesc    = `Edits todos.

  You can edit all facets of a todo.

  Read the full docs at https://ultodo.io/docs/cli/managing_tasks/#editing-todos`
		editCmdExample = `  To edit a todo's subject:
    ultodo edit 33 Meeting with @bob about +project
    ultodo e 33 Change the subject once again

  To edit just the due date, keeping the subject:
    ultodo edit 33 due:mon

  To remove a due date:
    ultodo edit 33 due none

  To edit a status
    ultodo edit 33 status:next

	To remove a status:
    ultodo edit 33 status:none`
	)

	var editCmd = &cobra.Command{
		Use:     "edit [id]",
		Aliases: []string{"e"},
		Example: editCmdExample,
		Long:    longDesc,
		Short:   editCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			todoID, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Could not parse todo ID: '%s'\n", args[0])
				return
			}
			ultodo.NewApp().EditTodo(todoID, strings.Join(args[1:], " "))
		},
	}

	rootCmd.AddCommand(editCmd)
}
