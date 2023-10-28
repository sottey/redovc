package cmd

import (
	"fmt"

	redovc "github.com/sottey/redo.vc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		editCmdDesc    = "Exports all todos"
		longDesc       = "Exports all todos"
		editCmdExample = `  Exports all todosTo export to CSV:
    redovc export csv

  To export to a text file
    redovc export txt`
	)

	var exportCmd = &cobra.Command{
		Use:     "export [json|csv|text]",
		Aliases: []string{"ex"},
		Example: editCmdExample,
		Long:    longDesc,
		Short:   editCmdDesc,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				fmt.Print("The export command requires a type argument. json, csv or text.\n")
				return
			}

			redovc.NewApp().Export(args[0])
		},
	}

	rootCmd.AddCommand(exportCmd)
}
