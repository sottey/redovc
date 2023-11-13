package cmd

import (
	"fmt"

	redovc "github.com/sottey/redo.vc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		serveCmdDesc    = "Serves todo data on specified ip:port location"
		longDesc        = "Serves todo data on specified ip:port location"
		serveCmdExample = `  Serves todo data on specified ip:port location
		Serves todo data on specified ip:port location:
    redovc serve 127.0.0.1:80

  To serve on localhost using port 8123
    redovc serve 127.0.0.1:8123`
	)

	var serveCmd = &cobra.Command{
		Use:     "serve [ip]:[port]",
		Aliases: []string{"srv"},
		Example: serveCmdExample,
		Long:    longDesc,
		Short:   serveCmdDesc,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				fmt.Print("The serve command requires an argument indicating the ip and port address. \n")
				return
			}

			redovc.NewApp().Serve(args[0])
		},
	}

	rootCmd.AddCommand(serveCmd)
}
