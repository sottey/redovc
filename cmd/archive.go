package cmd

import (
	"strings"

	redovc "github.com/sottey/redo.vc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		archiveCmdExample = `
  To archive a todo with id 33:
    redovc archive 33
    redovc ar 33

  To unarchive todo with id 33:
    redovc unarchive 33
    redovc uar 33

  To archive all completed todos:
    redovc archive completed
    redovc ar c

  Garbage collection will delete all archived todos, reclaming ids:
    redovc archive gc
    redovc ar gc

  redovc archive gc
  redovc ar gc
	  Run garbage collection. Delete all archived todos and reclaim ids`
	)

	var archiveCmd = &cobra.Command{
		Use:     `archive [id]`,
		Aliases: []string{"ar"},
		Example: archiveCmdExample,
		Short:   "Archives a todo.",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().ArchiveTodo(strings.Join(args, " "))
		},
	}

	var unarchiveCmd = &cobra.Command{
		Use:     "unarchive [id]",
		Aliases: []string{"uar"},
		Example: archiveCmdExample,
		Short:   "Un-archives a todo.",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().UnarchiveTodo(strings.Join(args, " "))
		},
	}

	var archiveCompletedCmd = &cobra.Command{
		Use:     "c",
		Example: "  redovc archive completed\n  redovc ar c",
		Short:   "Archives all completed todos.",
		Long:    `Archives all completed todos.`,
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().ArchiveCompleted()
		},
	}

	var archiveGarbageCollectCmd = &cobra.Command{
		Use:     "gc",
		Aliases: []string{"rm"},
		Short:   "Deletes all archived todos.",
		Long:    `Delete all archived todos, reclaiming ids.`,
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().GarbageCollect()
		},
	}

	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(unarchiveCmd)
	archiveCmd.AddCommand(archiveCompletedCmd)
	archiveCmd.AddCommand(archiveGarbageCollectCmd)
}
