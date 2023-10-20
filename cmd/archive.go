package cmd

import (
	"strings"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		archiveCmdExample = `
  To archive a todo with id 33:
    ultodo archive 33
    ultodo ar 33

  To unarchive todo with id 33:
    ultodo unarchive 33
    ultodo uar 33

  To archive all completed todos:
    ultodo archive completed
    ultodo ar c

  Garbage collection will delete all archived todos, reclaming ids:
    ultodo archive gc
    ultodo ar gc

  ultodo archive gc
  ultodo ar gc
	  Run garbage collection. Delete all archived todos and reclaim ids`
	)

	var archiveCmd = &cobra.Command{
		Use:     `archive [id]`,
		Aliases: []string{"ar"},
		Example: archiveCmdExample,
		Short:   "Archives a todo.",
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().ArchiveTodo(strings.Join(args, " "))
		},
	}

	var unarchiveCmd = &cobra.Command{
		Use:     "unarchive [id]",
		Aliases: []string{"uar"},
		Example: archiveCmdExample,
		Short:   "Un-archives a todo.",
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().UnarchiveTodo(strings.Join(args, " "))
		},
	}

	var archiveCompletedCmd = &cobra.Command{
		Use:     "c",
		Example: "  ultodo archive completed\n  ultodo ar c",
		Short:   "Archives all completed todos.",
		Long:    `Archives all completed todos.`,
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().ArchiveCompleted()
		},
	}

	var archiveGarbageCollectCmd = &cobra.Command{
		Use:     "gc",
		Aliases: []string{"rm"},
		Short:   "Deletes all archived todos.",
		Long:    `Delete all archived todos, reclaiming ids.`,
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().GarbageCollect()
		},
	}

	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(unarchiveCmd)
	archiveCmd.AddCommand(archiveCompletedCmd)
	archiveCmd.AddCommand(archiveGarbageCollectCmd)
}
