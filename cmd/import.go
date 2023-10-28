package cmd

import (
	"fmt"

	"github.com/sottey/redovc/redovc"
	"github.com/spf13/cobra"
)

func init() {
	var (
		importCmdDesc    = "Imports todos"
		longDesc         = "Inports todos"
		importCmdExample = `  Imports a list of todos. Note that there are specific formats required. See https://redovc/cmd/import.md for details
	Imports todos from a CSV file:
    redovc import ./import.csv

  To import to a json file
    redovc import txt`
	)

	var importCmd = &cobra.Command{
		Use:     "import filename",
		Aliases: []string{"im"},
		Example: importCmdExample,
		Long:    longDesc,
		Short:   importCmdDesc,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				fmt.Print("The import command requires a file argument which is the import file location and filename. \n")
				return
			}

			redovc.NewApp().Import(args[0])
		},
	}

	rootCmd.AddCommand(importCmd)
}
