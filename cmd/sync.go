package cmd

import (
	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		syncCmdDesc    = "Sync a list with ultodo.io"
		syncCmdExample = `  To synchronize your list:
    ultodo sync

  To set up your list to sync with ultodo.io:
    ultodo sync --setup

  To stop syncing your list with ultodo.io:
    ultodo sync --unsync

  Perform a sync without showing output to the screen:
    ultodo sync --quiet
	`

		syncCmdQuiet    bool
		setupCmd        bool
		unsyncCmd       bool
		syncCmdLongDesc = `Sync a list with ultodo.io.
  If you're using ultodo Pro, this will manually (and bi-directionally) sync your list with the remote list on ultodo.io.

  Note that you won't normally have to run this command.  Syncing occurs automatically when you manipulate your list locally.

  For more info on syncing, see https://ultodo.io/docs/cli/pro_integration`
	)

	var syncCmd = &cobra.Command{
		Use:     "sync",
		Example: syncCmdExample,
		Long:    syncCmdLongDesc,
		Short:   syncCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if setupCmd {
				ultodo.NewApp().SetupSync()
				return
			}
			if unsyncCmd {
				ultodo.NewApp().Unsync()
				return
			}

			ultodo.NewApp().Sync(syncCmdQuiet)
		},
	}

	syncCmd.Flags().BoolVarP(&syncCmdQuiet, "quiet", "q", false, "Run without output")
	syncCmd.Flags().BoolVarP(&setupCmd, "setup", "", false, "Set up a list to sync with ultodo.io, or pull a remote list to local")
	syncCmd.Flags().BoolVarP(&unsyncCmd, "unsync", "", false, "Stop syncing a list with ultodo.io")
	rootCmd.AddCommand(syncCmd)
}
