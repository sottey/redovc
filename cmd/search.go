package cmd

import (
	"strings"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		unicodeSupport    bool
		colorSupport      bool
		listNotes         bool
		showStatus        bool
		searchCmdDesc     = "Search todos"
		searchCmdExample  = "ultodo search:'this is the search string'"
		searchCmdLongDesc = "List todos that contain the specified string"
	)

	var searchCmd = &cobra.Command{
		Use:     "search",
		Aliases: []string{"sr"},
		Example: searchCmdExample,
		Long:    searchCmdLongDesc,
		Short:   searchCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewAppWithPrintOptions(unicodeSupport, colorSupport).ListTodos(strings.Join(args, " "), listNotes, showStatus)
		},
	}

	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&unicodeSupport, "unicode", "", true, "Allows unicode support in ultodo output")
	searchCmd.Flags().BoolVarP(&colorSupport, "color", "", true, "Allows color in ultodo output")
	searchCmd.Flags().BoolVarP(&listNotes, "notes", "", false, "Show a todo's notes when listing.")
	searchCmd.Flags().BoolVarP(&showStatus, "status", "", false, "Show a todo's status")
}
