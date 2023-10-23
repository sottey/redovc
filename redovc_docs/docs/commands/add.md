# add command
Adds a todo

  You can optionally specify a due date.
  This can be done by by putting 'due:[date]' at the end, where [date] is in:
  where [date] is one of:

`
tod - Items due today
today - Items due today
tom - Items due tomorrow
tomorrow - Items due tomorrow
thisweek - Items from the closest Monday to the following Sunday
nextweek - Items from next Monday through the following Sunday
lastweek - Items from Last Monday through the following Sunday
mon - Next Monday
tue - Next Tuesday
wed - Next Wednesday
thu - Next Thursday
fri - Next Friday
sat - Next Saturday
sun - Next Sunday
none - No due date
eocw - End of cal week move to next sunday
eow - End of week (Next sunday)
eoy - End of year
eoww - End of work week (Next Friday)
sow - Start of week (next Monday)
soww - Start of workweek (next Monday)
soy - Start Of year (next jan1)
[specific date] - Examples: oct11, 15jan, may30
`

  Dates can also be explicit, using 3 characters for the month.  They can be written in 2 different formats

`redovc a buy flowers for mom due:may12`

`redovc get halloween candy due:31oct`

  Todos can also recur.  Set the 'recur' directive to control recurrence

`redovc a Daily standup recur:weekdays`

`redovc a 1on1 meeting with jim recur:weekly`

## Usage
`redovc add <todo> [flags]`

## Aliases
  add, a

## Examples
`redovc add Prepare meeting notes about +importantProject for the meeting with #bob due:today`

`redovc add Meeting with #bob about +project due:tod`
  
`redovc a +work +verify did #john fix the build? due:tom`
  
`redovc a here is an important task priority:true recur:weekdays due:tom`

## Flags
`-h, --help   help for add`