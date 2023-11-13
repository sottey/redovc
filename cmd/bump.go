package cmd

import (
	redovc "github.com/sottey/redo.vc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		bumpCmdDesc    = "Bumps all overdue todos to be due today"
		longDesc       = "Bumps all overdue todos to be due today"
		bumpCmdExample = `  bump`
	)

	var bumpCmd = &cobra.Command{
		Use:     "bump",
		Example: bumpCmdExample,
		Long:    longDesc,
		Short:   bumpCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().Bump()
		},
	}

	rootCmd.AddCommand(bumpCmd)
}
