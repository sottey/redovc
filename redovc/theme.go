package redovc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type Theme struct {
	Columns []struct {
		Desc  string `json:"desc"`
		Name  string `json:"columnname"`
		Index int    `json:"index"`
	} `json:"Columns"`
	Colors []struct {
		Desc  string `json:"desc"`
		Name  string `json:"name"`
		Color string `json:"color"`
		Bold  bool   `json:"bold"`
	} `json:"Colors"`
}

// Color Objects
var (
	blackBoldFg     = color.New(color.Bold, color.FgBlack)
	blackFg         = color.New(0, color.FgBlack)
	blueBoldFg      = color.New(color.Bold, color.FgBlue)
	blueFg          = color.New(0, color.FgBlue)
	cyanBoldFg      = color.New(color.Bold, color.FgCyan)
	cyanFg          = color.New(0, color.FgCyan)
	greenBoldFg     = color.New(color.Bold, color.FgGreen)
	greenFg         = color.New(0, color.FgGreen)
	greyBoldFg      = color.New(color.Bold, color.FgHiBlack)
	greyFg          = color.New(0, color.FgHiBlack)
	hiBlueBoldFg    = color.New(color.Bold, color.FgHiBlue)
	hiBlueFg        = color.New(0, color.FgHiBlue)
	hiCyalBoldFg    = color.New(color.Bold, color.FgHiCyan)
	hiCyanFg        = color.New(0, color.FgHiCyan)
	hiGreenBoldFg   = color.New(color.Bold, color.FgHiGreen)
	hiGreenFg       = color.New(0, color.FgHiGreen)
	hiMagentaBoldFg = color.New(color.Bold, color.FgHiMagenta)
	hiMagentaFg     = color.New(0, color.FgHiMagenta)
	hiRedBoldFg     = color.New(color.Bold, color.FgHiRed)
	hiRedFg         = color.New(0, color.FgHiRed)
	hiWhiteBoldFg   = color.New(color.Bold, color.FgHiWhite)
	hiWhiteFg       = color.New(0, color.FgHiWhite)
	hiYellowBoldFg  = color.New(color.Bold, color.FgHiYellow)
	hiYellowFg      = color.New(0, color.FgHiYellow)
	magentaBoldFg   = color.New(color.Bold, color.FgMagenta)
	magentaFg       = color.New(0, color.FgMagenta)
	redBoldFg       = color.New(color.Bold, color.FgRed)
	redFg           = color.New(0, color.FgRed)
	whiteBoldFg     = color.New(color.Bold, color.FgWhite)
	whiteFg         = color.New(0, color.FgWhite)
	yellowBoldFg    = color.New(color.Bold, color.FgYellow)
	yellowFg        = color.New(0, color.FgYellow)
)

// Default String Type Colors
var (
	groupTitleColor             = cyanFg
	noteIDColor                 = cyanFg
	noteTextColor               = whiteFg
	taskIDColor                 = yellowFg
	taskIDPriColor              = yellowBoldFg
	completedColor              = whiteFg
	statusColor                 = greenFg
	statusPriColor              = greenBoldFg
	informationColor            = whiteFg
	todayColor                  = greenFg
	todayPriColor               = greenBoldFg
	tomorrowColor               = yellowFg
	tomorrowPriColor            = yellowBoldFg
	overdueColor                = redFg
	overduePriColor             = redBoldFg
	otherDue                    = cyanFg
	otherDuePriColor            = cyanBoldFg
	taskTextColor               = whiteFg
	taskTextProjectWordColor    = magentaFg
	taskTextTagWordColor        = greenFg
	taskTextPriColor            = whiteBoldFg
	taskTextProjectPriWordColor = magentaBoldFg
	taskTextTagWordPriColor     = greenBoldFg
)

func LoadTheme() {
	theme, err := GetTheme()
	if err != nil {
		fmt.Printf("Theme file found but could not load: %v\n", err)
	}

	for _, color := range theme.Colors {

		name := strings.ToLower(color.Name)
		color := StringToColor(color.Color, color.Bold)

		SetThemeColor(name, color)
	}
}

// Get the themes configuration file
func GetTheme() (Theme, error) {
	var theme Theme

	themeFile := GetLocation()
	data, err := os.ReadFile(themeFile)
	if err != nil {
		return theme, nil
	}

	jerr := json.Unmarshal(data, &theme)
	if jerr != nil {
		return theme, jerr
	}

	return theme, nil
}

// GetLocation is returning the location of the .redovc.theme.json file.
func GetLocation() string {
	if LocalThemeFileExists() {
		dir, _ := os.Getwd()
		localrepo := filepath.Join(dir, TodosThemeFile)
		return localrepo
	}
	return fmt.Sprintf("%s/%s", UserDataDir(), TodosThemeFile)
}

// Returns if a local .redovc.theme.json file exists in the current dir.
func LocalThemeFileExists() bool {
	dir, _ := os.Getwd()
	localrepo := filepath.Join(dir, TodosThemeFile)
	_, err := os.Stat(localrepo)
	return err == nil
}

// Set a string target to theme color
func SetThemeColor(name string, color *color.Color) {

	switch name {
	case "grouptitlecolor":
		groupTitleColor = color
	case "noteidcolor":
		noteIDColor = color
	case "notetextcolor":
		noteTextColor = color
	case "taskidcolor":
		taskIDColor = color
	case "taskidpricolor":
		taskIDPriColor = color
	case "completedcolor":
		completedColor = color
	case "statuscolor":
		statusColor = color
	case "statuspricolor":
		statusPriColor = color
	case "informationcolor":
		informationColor = color
	case "todaycolor":
		todayColor = color
	case "todaypricolor":
		todayPriColor = color
	case "tomorrowcolor":
		tomorrowColor = color
	case "tomorrowpricolor":
		tomorrowPriColor = color
	case "overduecolor":
		overdueColor = color
	case "overduepricolor":
		overduePriColor = color
	case "otherdue":
		otherDue = color
	case "otherduepricolor":
		otherDuePriColor = color
	case "tasktextcolor":
		taskTextColor = color
	case "tasktextprojectwordcolor":
		taskTextProjectWordColor = color
	case "tasktexttagwordcolor":
		taskTextTagWordColor = color
	case "tasktextpricolor":
		taskTextPriColor = color
	case "tasktextprojectpriwordcolor":
		taskTextProjectPriWordColor = color
	case "tasktexttagwordpricolor":
		taskTextTagWordPriColor = color
	}
}

// Convert a color string to a color object
func StringToColor(colorString string, bold bool) *color.Color {
	colorString = strings.ToLower(colorString)

	switch colorString {
	case "blue":
		if !bold {
			return blueFg
		} else {
			return blueBoldFg
		}
	case "green":
		if !bold {
			return greenFg
		} else {
			return greenBoldFg
		}
	case "cyan":
		if !bold {
			return cyanFg
		} else {
			return cyanBoldFg
		}
	case "magenta":
		if !bold {
			return magentaFg
		} else {
			return magentaBoldFg
		}
	case "red":
		if !bold {
			return redFg
		} else {
			return redBoldFg
		}
	case "white":
		if !bold {
			return whiteFg
		} else {
			return whiteBoldFg
		}
	case "yellow":
		if !bold {
			return yellowFg
		} else {
			return yellowBoldFg
		}
	case "black":
		if !bold {
			return blackFg
		} else {
			return blackBoldFg
		}
	case "grey":
		if !bold {
			return greyFg
		} else {
			return greyBoldFg
		}
	case "hiblue":
		if !bold {
			return hiBlueFg
		} else {
			return hiBlueBoldFg
		}
	case "hicyan":
		if !bold {
			return hiCyanFg
		} else {
			return hiCyalBoldFg
		}
	case "higreen":
		if !bold {
			return hiGreenFg
		} else {
			return hiGreenBoldFg
		}
	case "himagenta":
		if !bold {
			return hiMagentaFg
		} else {
			return hiMagentaBoldFg
		}
	case "hired":
		if !bold {
			return hiRedFg
		} else {
			return hiRedBoldFg
		}
	case "hiwhite":
		if !bold {
			return hiWhiteFg
		} else {
			return hiWhiteBoldFg
		}
	case "hiyellow":
		if !bold {
			return hiYellowFg
		} else {
			return hiYellowBoldFg
		}
	default:
		if !bold {
			return whiteFg
		} else {
			return whiteBoldFg
		}
	}
}

func OrderColumns(showStatus bool, id string, completed string, info string, due string, status string, subject string) []string {
	statusIndex := -1
	columnList := []string{"", "", "", "", "", ""}
	theme, err := GetTheme()
	if err != nil {
		fmt.Printf("Theme file found but could not load: %v\n", err)
	}

	for _, column := range theme.Columns {
		colName := column.Name
		index := column.Index

		switch colName {
		case "id":
			columnList[index] = id
		case "completed":
			columnList[index] = completed
		case "information":
			columnList[index] = info
		case "due":
			columnList[index] = due
		case "status":
			statusIndex = index
			columnList[index] = status
		case "subject":
			columnList[index] = subject
		}
	}

	if !showStatus {
		RemoveFromStringArray(columnList, statusIndex)
	}

	return columnList
}

const themeTemplate = `
{
    "Colors":
    [
        {
            "desc": "Color of group names when group: is specified",
            "name": "groupTitleColor",
            "color": "cyan",
            "bold": true
        },
        {
            "desc":"Color of note id's when --notes is used",
            "name": "noteIDColor",
            "color": "white",
            "bold": false
        },
        {
            "desc":"Color of note contents when --notes is used",
            "name": "noteTextColor",
            "color": "white",
            "bold": false
        },
        {
            "desc":"Color of a task id",
            "name": "taskIDColor",
            "color": "yellow",
            "bold": false
        },
        {
            "desc":"Color of a task id if the task is prioritized",
            "name": "taskIDPriColor",
            "color": "yellow",
            "bold": true
        },
        {
            "desc":"Color of completed indicator ([ ] or [X])",
            "name": "completedColor",
            "color": "white",
            "bold": false
        },
        {
            "desc":"Color of task status",
            "name": "statusColor",
            "color": "green",
            "bold": false
        },
        {
            "desc":"Color of a task status if task is prioritized",
            "name": "statusPriColor",
            "color": "green",
            "bold": true
        },
        {
            "desc":"Color of area indicating Priority, Note and Archived",
            "name": "informationColor",
            "color": "white",
            "bold": false
        },
        {
            "desc":"Color of a task due today",
            "name": "todayColor",
            "color": "green",
            "bold": false
        },
        {
            "desc":"Color of a task due today if that task is prioritized",
            "name": "todayPriColor",
            "color": "green",
            "bold": true
        },
        {
            "desc":"Color of a task due tomorrow",
            "name": "tomorrowColor",
            "color": "yellow",
            "bold": false
        },
        {
            "desc":"Color of a task due tomorrow if that task is prioritized",
            "name": "tomorrowPriColor",
            "color": "yellow",
            "bold": true
        },
        {
            "desc":"Color of a task that is overdue",
            "name": "overdueColor",
            "color": "red",
            "bold": false
        },
        {
            "desc":"Color of a task that is overdue if that task is prioritized",
            "name": "overduePriColor",
            "color": "red",
            "bold": true
        },
        {
            "desc":"Color of a task that is not today, tomorrow or overdue",
            "name": "otherDue",
            "color": "cyan",
            "bold": false
        },
        {
            "desc":"Color of a task that is not today, tomorrow or overdue if the task is prioritized",
            "name": "otherDuePriColor",
            "color": "cyan",
            "bold": true
        },
        {
            "desc":"Color of task text",
            "name": "taskTextColor",
            "color": "white",
            "bold": false
        },
        {
            "desc":"Color of project names in task text",
            "name": "taskTextProjectWordColor",
            "color": "magents",
            "bold": false
        },
        {
            "desc":"Color of tags in task text",
            "name": "taskTextTagWordColor",
            "color": "green",
            "bold": false
        },
        {
            "desc":"Color of task text if the task is prioritized",
            "name": "taskTextPriColor",
            "color": "white",
            "bold": true
        },
        {
            "desc":"Color of project names in task text if that task is prioritized",
            "name": "taskTextProjectPriWordColor",
            "color": "white",
            "bold": true
        },
        {
            "desc":"Color of tags in task text if that task is prioritized",
            "name": "taskTextTagWordPriColor",
            "color": "green",
            "bold": true
        }
    ],
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
            "desc": "Task subject - Setting this to something other than index 5 could really mess up the screen",
            "columnname": "subject",
            "index":5
        }
    ]
}
	`
