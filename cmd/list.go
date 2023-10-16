package cmd

import (
	"strings"

	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		unicodeSupport bool
		colorSupport   bool
		listNotes      bool
		showStatus     bool
		listCmdDesc    = "List todos."
		listCmdExample = `
  Filtering by date:
  ------------------

  ultodo list due:<date>
  ultodo list duebefore:<date>
  ultodo list dueafter:<date>

  where <date> is one of:
  (tod|today|tom|tomorrow|thisweek|nextweek|lastweek|mon|tue|wed|thu|fri|sat|sun|none|<specific date>)

  List all todos due today:
    ultodo list due:tod

  Lists all todos due tomorrow:
    ultodo list due:tom

  Lists all todos due monday:
    ultodo list due:mon

  Lists all todos with no due date:
    ultodo list due:none

  Lists all overdue todos:
    ultodo list duebefore:today

  Lists all todos in due in the future:
    ultodo list dueafter:today

  When using a specific date, it needs to be in the format of jun23 or 23jun:
    ultodo list due:jun23

  Filtering by status:
  --------------------

  List all todos with a status of "started"
    ultodo list status:started

  List all todos without a status of "started"
    ultodo list status:-started

  List all todos without a status of "started" or "finished"
    ultodo list status:-started,-finished

  Filtering by projects or contexts:
  ----------------------------------

  Project and context filtering are very similar:
    ultodo list project:<project>
    ultodo list context:<context>

  List all todos with a project of "mobile"
    ultodo list project:mobile

  List all todos with a project of "mobile" and "devops"
    ultodo list project:mobile,devops

  List all todos with a project of "mobile" but not "devops"
    ultodo list project:mobile,-devops

  List all todos without a project of "devops"
    ultodo list project:-devops

  Filtering by priority, completed, etc:
  --------------------------------------

  You can filter todos on their priority or completed status:
    ultodo list is:priority
    ultodo list not:priority

    ultodo list is:completed
    ultodo list not:completed

  There are additional filters for showing completed todos:
    ultodo list completed:today
    ultodo list completed:thisweek

  By default, ultodo will not show archived todos. To show archived todos:
    ultodo list is:archived

  Grouping:
  ---------

  List all todos grouped by context:
    ultodo list group:c

  List all todos grouped by project:
    ultodo list group:p

  List all todos grouped by status:
 	  ultodo list group:s

  Combining filters:
  ------------------

  Of course, you can combine grouping and filtering to get a nice formatted list.

  Lists all todos due today grouped by context:
    ultodo list group:c due:today

  Lists all todos due today for +mobile, grouped by context:
    ultodo list project:mobile group:c due:thisweek

  Lists all prioritized todos that are not completed and are overdue.  Include a todo's notes when listing:
    ultodo list --notes is:priority not:completed duebefore:tod

  Lists all todos due tomorrow concerning @frank for +project, grouped by project:
    ultodo list context:frank group:p due:tom

  Indicator flags
  ---------------

  If you pass --status=true as a flag, you'll see an extra column when listing todos.

  * = Todo is prioritized
  N = Todo has notes attached
  A = Todo is archived
`
		listCmdLongDesc = `List todos, optionally providing a filter.

When listing todos, you can apply powerful filters, and perform grouping.

See the full docs at https://ultodo.io/docs/cli/showing_tasks`
	)

	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "ls"},
		Example: listCmdExample,
		Long:    listCmdLongDesc,
		Short:   listCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewAppWithPrintOptions(unicodeSupport, colorSupport).ListTodos(strings.Join(args, " "), listNotes, showStatus)
		},
	}

	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&unicodeSupport, "unicode", "", true, "Allows unicode support in ultodo output")
	listCmd.Flags().BoolVarP(&colorSupport, "color", "", true, "Allows color in ultodo output")
	listCmd.Flags().BoolVarP(&listNotes, "notes", "", false, "Show a todo's notes when listing. ")
	listCmd.Flags().BoolVarP(&showStatus, "status", "", false, "Show a todo's status")
}
