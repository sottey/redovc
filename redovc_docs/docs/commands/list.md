# list command
List todos, optionally providing a filter

When listing todos, you can apply powerful filters, and perform grouping.

## Usage

`redovc list [flags]`

## Aliases
  list, l, ls

## Examples

#### Filtering by date:
  
`redovc list due:[date]`

`redovc list duebefore:[date]`

`redovc list dueafter:[date]`

where [date] is one of:

`
tod
today
tom
tomorrow
thisweek
nextweek
lastweek
mon
tue
wed
thu
fri
sat
sun
none
[specific date]
`


##### List all todos due today:

`redovc list due:tod`

##### Lists all todos due tomorrow:

`redovc list due:tom`

##### Lists all todos due monday:

`redovc list due:mon`

##### Lists all todos with no due date:

`redovc list due:none`

##### Lists all overdue todos:

`redovc list duebefore:today`

##### Lists all todos in due in the future:

`redovc list dueafter:today`

##### When using a specific date, it needs to be in the format of jun23 or 23jun:

`redovc list due:jun23`

#### Filtering by status:

##### List all todos with a status of "started"

`redovc list status:started`

##### List all todos without a status of "started"

`redovc list status:-started`

##### List all todos without a status of "started" or "finished"

`redovc list status:-started,-finished`

#### Filtering by projects or tags:

##### Project and tag filtering are very similar:

`redovc list project:[project]`

`redovc list tag:[tag]`

##### List all todos with a project of "mobile"

`redovc list project:mobile`

##### List all todos with a project of "mobile" and "devops"

`redovc list project:mobile,devops`

##### List all todos with a project of "mobile" but not "devops"

`redovc list project:mobile,-devops`

##### List all todos without a project of "devops"

`redovc list project:-devops`

#### Filtering by priority, completed, etc:

##### You can filter todos on their priority or completed status:

`redovc list is:priority`

`redovc list not:priority`

`redovc list is:completed`

`redovc list not:completed`

##### There are additional filters for showing completed todos:

`redovc list completed:today`

`redovc list completed:thisweek`

##### By default, redovc will not show archived todos. To show archived todos:

`redovc list is:archived`

#### Grouping:

##### List all todos grouped by tag:

`redovc list group:t`

##### List all todos grouped by project:

`redovc list group:p`

##### List all todos grouped by status:

`redovc list group:s`

#### Combining filters:

Of course, you can combine grouping and filtering to get a nice formatted list.

##### Lists all todos due today grouped by tag:

`redovc list group:t due:today`

##### Lists all todos due today for +mobile, grouped by tag:

`redovc list project:mobile group:t due:thisweek`

##### Lists all prioritized todos that are not completed and are overdue.  Include a todo's notes when listing:

`redovc list --notes is:priority not:completed duebefore:tod`

##### Lists all todos due tomorrow concerning #frank for +project, grouped by project:

`redovc list tag:frank group:p due:tom`

##### Indicator flags

If you pass --status=true as a flag, you'll see an extra column when listing todos.

###### P = Todo is prioritized
###### N = Todo has notes attached
###### A = Todo is archived


## Flags

`--color     Allows color in redovc output (default: true) (default true)`

`-h, --help  Help for list`

`--notes     Show a todo's notes when listing. (default: false)`

`--status    Show a todo's status (default: false)`

`--unicode   Allows unicode support in redovc output (default: true) (default true)`