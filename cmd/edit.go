package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sottey/redo.vc/redovc"
	"github.com/spf13/cobra"
)

func init() {
	var (
		editCmdDesc = "Edits todos"
		longDesc    = `Edits todos.

  You can edit all facets of a todo.`
		editCmdExample = `  To edit a todo's subject:
    redovc edit 33 Meeting with #bob about +project
    redovc e 33 Change the subject once again

  To edit just the due date, keeping the subject:
    redovc edit 33 due:mon

  To remove a due date:
    redovc edit 33 due none

  To edit a status
    redovc edit 33 status:next

	To remove a status:
    redovc edit 33 status:none`
	)

	var editCmd = &cobra.Command{
		Use:     "edit [id]",
		Aliases: []string{"e"},
		Example: editCmdExample,
		Long:    longDesc,
		Short:   editCmdDesc,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				fmt.Print("The edit command requires arguments. See ultralist e[dit] -h.\n")
				return
			}

			todoID, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Could not parse todo ID: '%s'\n", args[0])
				return
			}
			redovc.NewApp().EditTodo(todoID, strings.Join(args[1:], " "))
		},
	}

	rootCmd.AddCommand(editCmd)
}
