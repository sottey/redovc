package cmd

import (
	"fmt"
	"strconv"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		long = `Delete a note from a todo.
  For more info, see https://ultodo.io/docs/cli/managing_tasks/#notes-management`
		example = `  To see your todos with notes:
    ultodo list --notes

  To delete note 0 from todo 3:
    ultodo dn 3 0`
	)

	var deleteNoteCmd = &cobra.Command{
		Use:     "deletenote <todoID> <noteID>",
		Aliases: []string{"dn"},
		Long:    long,
		Example: example,
		Short:   "Delete a note from a todo.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 2 {
				todoID, _ := strconv.Atoi(args[0])
				noteID, _ := strconv.Atoi(args[1])
				ultodo.NewApp().DeleteNote(todoID, noteID)
			} else {
				fmt.Printf("todoID and noteID not specified\n")
			}
		},
	}

	rootCmd.AddCommand(deleteNoteCmd)
}
