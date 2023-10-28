package cmd

import (
	"fmt"

	redovc "github.com/sottey/redo.vc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		versionCmdDesc     = "Displays the version of redovc"
		versionCmdLongDesc = versionCmdDesc + "."
	)

	var versionCmd = &cobra.Command{
		Use:   "version",
		Long:  versionCmdLongDesc,
		Short: versionCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("redovc v%s\n", redovc.VERSION)
		},
	}

	rootCmd.AddCommand(versionCmd)
}
