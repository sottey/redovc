package cmd

import (
	"strconv"
	"strings"

	"github.com/sottey/redo.vc/redovc"
	"github.com/spf13/cobra"
)

func init() {
	var (
		cmdDesc     = "Edits notes on a todo."
		longCmdDesc = "Edits notes on a todo."
		example     = `  To see your todos with notes:
    redovc list --notes

  To edit note 0 from todo 3:
    redovc en 3 0 this is the new note`
	)
	var editNoteCmd = &cobra.Command{
		Use:     "editnote",
		Aliases: []string{"en"},
		Example: example,
		Long:    longCmdDesc,
		Short:   cmdDesc,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			todoID, _ := strconv.Atoi(args[0])
			noteID, _ := strconv.Atoi(args[1])
			redovc.NewApp().EditNote(todoID, noteID, strings.Join(args[2:], " "))
		},
	}

	rootCmd.AddCommand(editNoteCmd)
}
