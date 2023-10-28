package cmd

import (
	"strconv"
	"strings"

	redovc "github.com/sottey/redo.vc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		addNoteCmdDesc    = "Adds notes to a todo."
		addNoteCmdExample = "  redovc an 1 this is a note for the first todo"
	)

	var addNoteCmd = &cobra.Command{
		Use:     "addnote <todoID> <noteContent>",
		Aliases: []string{"an"},
		Example: addNoteCmdExample,
		Short:   addNoteCmdDesc,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			todoID, _ := strconv.Atoi(args[0])
			redovc.NewApp().AddNote(todoID, strings.Join(args[1:], " "))
		},
	}
	rootCmd.AddCommand(addNoteCmd)
}
