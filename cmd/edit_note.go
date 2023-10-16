package cmd

import (
	"strconv"
	"strings"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		cmdDesc     = "Edits notes on a todo."
		longCmdDesc = "Edits notes on a todo.\n For more info, see https://ultodo.io/docs/cli/managing_tasks/#notes-management"
		example     = `  To see your todos with notes:
    ultodo list --notes

  To edit note 0 from todo 3:
    ultodo en 3 0 this is the new note`
	)
	var editNoteCmd = &cobra.Command{
		Use:     "editnote",
		Aliases: []string{"en"},
		Example: example,
		Long:    longCmdDesc,
		Short:   cmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			todoID, _ := strconv.Atoi(args[0])
			noteID, _ := strconv.Atoi(args[1])
			ultodo.NewApp().EditNote(todoID, noteID, strings.Join(args[2:], " "))
		},
	}

	rootCmd.AddCommand(editNoteCmd)
}
