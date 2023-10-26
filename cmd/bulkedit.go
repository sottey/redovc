package cmd

import (
	"github.com/sottey/redo.vc/redovc"
	"github.com/spf13/cobra"
)

func init() {
	var (
		bulkeditCmdExample = `To edit the .todos.json file:
    redovc bulkedit
    redovc be`
	)

	var bulkeditCmd = &cobra.Command{
		Use:     `bulkedit`,
		Aliases: []string{"be"},
		Example: bulkeditCmdExample,
		Long:    "bulkedit will open either the .todos.json file in the current directory or, if missing, the .todos.json file in the home directory",
		Short:   "Edit raw .todos.json file",
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().BulkEdit()
		},
	}

	rootCmd.AddCommand(bulkeditCmd)
}
