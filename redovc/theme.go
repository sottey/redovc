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
		Desc    string `json:"desc"`
		Nameame string `json:"columnname"`
		Index   int    `json:"index"`
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
	blueFg        = color.New(0, color.FgBlue)
	blueBoldFg    = color.New(color.Bold, color.FgBlue)
	greenFg       = color.New(0, color.FgGreen)
	greenBoldFg   = color.New(color.Bold, color.FgGreen)
	cyanFg        = color.New(0, color.FgCyan)
	cyanBoldFg    = color.New(color.Bold, color.FgCyan)
	magentaFg     = color.New(0, color.FgMagenta)
	magentaBoldFg = color.New(color.Bold, color.FgMagenta)
	redFg         = color.New(0, color.FgRed)
	redBoldFg     = color.New(color.Bold, color.FgRed)
	whiteFg       = color.New(0, color.FgWhite)
	whiteBoldFg   = color.New(color.Bold, color.FgWhite)
	yellowFg      = color.New(0, color.FgYellow)
	yellowBoldFg  = color.New(color.Bold, color.FgYellow)
	blackFg       = color.New(0, color.FgBlack)
	blackBoldFg   = color.New(color.Bold, color.FgBlack)
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

// GetLocation is returning the location of the .todos.theme.json file.
func GetLocation() string {
	if LocalThemeFileExists() {
		dir, _ := os.Getwd()
		localrepo := filepath.Join(dir, TodosThemeFile)
		return localrepo
	}
	return fmt.Sprintf("%s/%s", UserHomeDir(), TodosThemeFile)
}

// Returns if a local .todos..theme.json file exists in the current dir.
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
	default:
		if !bold {
			return whiteFg
		} else {
			return whiteBoldFg
		}
	}
}

const themeTemplate = `
{
    "Columns":
    [
        {
            "desc": "Column1 Desc",
            "columnname": "column1",
            "index":0
        },
        {
            "desc": "Column2 Desc",
            "columnname": "column2",
            "index":1
        }
    ],

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
    ]
}
	`
