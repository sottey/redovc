package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var (
		webCmdDesc     = "Open your list on ultodo.io"
		webCmdLongDesc = "\nIf your list is synced with ultodo.io, use this command to open your list with your web browser."
	)

	var webCmd = &cobra.Command{
		Use:   "web",
		Long:  webCmdLongDesc,
		Short: webCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Web is not supported. For web syncing, see ultodo.io\n")
		},
	}

	rootCmd.AddCommand(webCmd)
}
