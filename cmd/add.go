package cmd

import (
	"strings"

	redovc "github.com/sottey/redovc/lib"
	"github.com/spf13/cobra"
)

func init() {
	var (
		addCmdDesc    = "Adds todos"
		addCmdExample = `  redovc add Prepare meeting notes about +importantProject for the meeting with #bob due:today
  redovc add Meeting with #bob about +project due:tod
  redovc a +work +verify did #john fix the build? due:tom
  redovc a here is an important task priority:true recur:weekdays due:tom`

		addCmdLongDesc = `Adds todos.

  You can optionally specify a due date.
  This can be done by by putting 'due:<date>' at the end, where <date> is in (tod|today|tom|tomorrow|mon|tue|wed|thu|fri|sat|sun|thisweek|nextweek).

  Dates can also be explicit, using 3 characters for the month.  They can be written in 2 different formats:
    redovc a buy flowers for mom due:may12
    redovc get halloween candy due:31oct

  Todos can also recur.  Set the 'recur' directive to control recurrence:
    redovc a Daily standup recur:weekdays
    redovc a 1on1 meeting with jim recur:weekly
`
	)

	var addCmd = &cobra.Command{
		Use:     "add <todo>",
		Aliases: []string{"a"},
		Example: addCmdExample,
		Long:    addCmdLongDesc,
		Short:   addCmdDesc,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().AddTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(addCmd)
}
