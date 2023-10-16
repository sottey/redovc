package cmd

import (
	"fmt"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		versionCmdDesc     = "Displays the version of ultodo"
		versionCmdLongDesc = versionCmdDesc + "."
	)

	var versionCmd = &cobra.Command{
		Use:   "version",
		Long:  versionCmdLongDesc,
		Short: versionCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ultodo v%s\n", ultodo.VERSION)
		},
	}

	rootCmd.AddCommand(versionCmd)
}
