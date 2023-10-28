package cmd

import (
	"strings"

	redovc "github.com/sottey/redovc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		deleteCmdDesc    = "Deletes todos"
		deleteCmdExample = `  To delete a todo with ID 33:
    redovc d 33
    redovc delete 33

  Note, this will also free up the id of 33.`
		deleteCmdLongDesc = `Delete a todo with a specified ID.`
	)

	var deleteCmd = &cobra.Command{
		Use:     "delete [id]",
		Aliases: []string{"d", "rm"},
		Example: deleteCmdExample,
		Long:    deleteCmdLongDesc,
		Short:   deleteCmdDesc,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().DeleteTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(deleteCmd)
}
