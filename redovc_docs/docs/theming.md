# Theming

## Themes
Command shells have very limited color abilities, but I decided it would be good to not impose my color scheme on everyone who uses redovc.... So, theming.

## Theme file
A theme file can live in your home directory and it will be used as the default. You can also have theme files in other directories and when you are in that directory, that theme will be used.

The theme file must be named .todos.theme.json and the contents are used to describe how different sections of the redovc list command presents data.

## Theme file color structure
You can find sample theme files in the [sampleThemes](https://github.com/sottey/redo.vc/tree/main/sampleThemes) directory of the source code.

For each section, there are a number of properties:

- Desc - Color of group names when group: is specified
- Name - Name of section. These are used by the code and must be one of:
    - groupTitleColor - Color of group names when group: is specified
    - noteIDColor - Color of note id's when --notes is used
    - noteTextColor - Color of note contents when --notes is used
    - taskIDColor - Color of a task id
    - taskIDPriColor - Color of a task id if the task is prioritized
    - completedColor - Color of completed indicator ([ ] or [X])
    - statusColor - Color of task status
    - statusPriColor - Color of a task status if task is prioritized
    - informationColor - Color of area indicating Priority, Note and Archived
    - todayColor - Color of a task due today
    - todayPriColor - Color of a task due today if that task is prioritized
    - tomorrowColor - Color of a task due tomorrow
    - tomorrowPriColor - Color of a task due tomorrow if that task is prioritized
    - overdueColor - Color of a task that is overdue
    - overduePriColor - Color of a task that is overdue if that task is prioritized
    - otherDue - Color of a task that is not today, tomorrow or overdue
    - otherDuePriColor - Color of a task that is not today, tomorrow or overdue if the task is prioritized
    - taskTextColor - Color of task text
    - taskTextProjectWordColor - Color of project names in task text
    - taskTextTagWordColor - Color of tags in task text
    - taskTextPriColor - Color of task text if the task is prioritized
    - taskTextProjectPriWordColor - Color of project names in task text if that task is prioritized
    - taskTextTagWordPriColor - Color of tags in task text if that task is prioritized
- Color - The color to be used. Valid colors are:
	- black
	- red
	- green
	- yellow
	- blue
	- magenta
	- cyan
	- white
- Bold = true or false

## Theme File Column Ordering Structure (not yet implemented):
I am working on people being able to specify what order the columns are displayed when using 'redovc list'. These will be contained in the theme file as well.

Currently, the json looks like this:
```
    "Columns":
    [
        {
            "desc": "Task ID",
            "columnname": "id",
            "index":0
        },
        {
            "desc": "Completed column",
            "columnname": "completed",
            "index":1
        },
        {
            "desc": "Information - Priority, Note and Archived flags",
            "columnname": "information",
            "index":2
        },
        {
            "desc": "Task Due date",
            "columnname": "due",
            "index":3
        },
        {
            "desc": "Task status",
            "columnname": "status",
            "index":4
        },
        {
            "desc": "Task subject - Changing this index to anything other than 5 will probably make everything look wonky.",
            "columnname": "subject",
            "index":5
        }
    ],
```