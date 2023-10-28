package redovc

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/fatih/color"
)

// ScreenPrinter is the default struct of this file
type ScreenPrinter struct {
	Writer         *io.Writer
	UnicodeSupport bool
}

const (
	idColumnWidth        = "%-4s"
	completedColumnWidth = "%-3s"
	infoColumnWidth      = "%-3s"
	dueColumnWidth       = "%-10s"
	statusColumnWidth    = "%-10s"
)

// NewScreenPrinter creates a new screeen printer.
func NewScreenPrinter(unicodeSupport bool) *ScreenPrinter {
	w := new(io.Writer)
	formatter := &ScreenPrinter{Writer: w, UnicodeSupport: unicodeSupport}
	return formatter
}

// Print prints the output of redovc to the terminal screen.
func (f *ScreenPrinter) Print(groupedTodos *GroupedTodos, printNotes bool, showStatus bool) {
	var keys []string
	for key := range groupedTodos.Groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	LoadTheme()

	tabby := tabby.NewCustom(tabwriter.NewWriter(color.Output, 0, 0, 2, ' ', 0))
	tabby.AddLine()
	for _, key := range keys {
		var title = key + " (" + strconv.Itoa(len(groupedTodos.Groups[key])) + ")"
		tabby.AddLine(groupTitleColor.Sprint(title))
		for _, todo := range groupedTodos.Groups[key] {
			f.printTodo(tabby, todo, printNotes, showStatus)
		}
		tabby.AddLine()
	}
	tabby.Print()
}

func (f *ScreenPrinter) printTodo(tabby *tabby.Tabby, todo *Todo, printNotes bool, showStatus bool) {
	id := fmt.Sprintf(idColumnWidth, f.formatID(todo.ID, todo.IsPriority))
	completed := fmt.Sprintf(completedColumnWidth, f.formatCompleted(todo.Completed))
	info := fmt.Sprintf(infoColumnWidth, f.formatInfoFlags(todo))
	due := fmt.Sprintf(dueColumnWidth, f.formatDue(todo.Due, todo.IsPriority, todo.Completed))
	status := fmt.Sprintf(statusColumnWidth, f.formatStatus(todo.Status, todo.IsPriority))
	subject := f.formatSubject(todo.Subject, todo.IsPriority)

	columns := OrderColumns(showStatus, id, completed, info, due, status, subject)

	if showStatus {
		tabby.AddLine(columns[0], columns[1], columns[2], columns[3], columns[4], columns[5])
	} else {
		tabby.AddLine(columns[0], columns[1], columns[2], columns[3], columns[4])
	}

	if printNotes {
		for nid, note := range todo.Notes {
			tabby.AddLine(
				"    "+noteIDColor.Sprint(strconv.Itoa(nid)),
				noteTextColor.Sprintf(idColumnWidth, ""),
				noteTextColor.Sprintf(completedColumnWidth, ""),
				noteTextColor.Sprintf(infoColumnWidth, ""),
				noteTextColor.Sprintf(dueColumnWidth, ""),
				noteTextColor.Sprintf(statusColumnWidth, ""),
				noteTextColor.Sprintf(""),
				noteTextColor.Sprint("- "+note))
		}
	}
}

func (f *ScreenPrinter) formatID(ID int, isPriority bool) string {
	if isPriority {
		return taskIDPriColor.Sprintf(strconv.Itoa(ID))
	}
	return taskIDColor.Sprintf(strconv.Itoa(ID))
}

func (f *ScreenPrinter) formatCompleted(completed bool) string {
	if completed {
		if f.UnicodeSupport {
			return completedColor.Add(color.CrossedOut).Sprint("[✔]")
		}
		return completedColor.Add(color.CrossedOut).Sprint("[x]")
	}
	return completedColor.Sprint("[ ]")
}

func (f *ScreenPrinter) formatDue(due string, isPriority bool, completed bool) string {
	if due == "" {
		return whiteFg.Sprintf("")
	}
	dueTime, _ := time.Parse(DATE_FORMAT, due)

	if isPriority {
		return f.printPriorityDue(dueTime, completed)
	}
	return f.printDue(dueTime, completed)
}

func (f *ScreenPrinter) formatStatus(status string, isPriority bool) string {
	if status == "" {
		return statusColor.Sprintf("")
	}

	//statusRune := []rune(status)

	if isPriority {
		return statusPriColor.Sprintf(status) //string(statusRune[0:10]))
	}
	return statusColor.Sprintf(status) //string(statusRune[0:10]))
}

func (f *ScreenPrinter) formatInfoFlags(todo *Todo) string {
	var infoFlags []string
	if todo.IsPriority {
		infoFlags = append(infoFlags, "P")
	} else {
		infoFlags = append(infoFlags, "-")
	}
	if todo.HasNotes() {
		infoFlags = append(infoFlags, "N")
	} else {
		infoFlags = append(infoFlags, "-")
	}
	if todo.Archived {
		infoFlags = append(infoFlags, "A")
	} else {
		infoFlags = append(infoFlags, "-")
	}

	return informationColor.Sprint(strings.Join(infoFlags, ""))
}

func (f *ScreenPrinter) printDue(due time.Time, completed bool) string {
	if isToday(due) {
		return todayColor.Sprintf("today")
	} else if isTomorrow(due) {
		return tomorrowColor.Sprintf("tomorrow")
	} else if isPastDue(due) && !completed {
		return overdueColor.Sprintf(due.Format("Mon Jan 02"))
	}

	return otherDue.Sprintf(due.Format("Mon Jan 02"))
}

func (f *ScreenPrinter) printPriorityDue(due time.Time, completed bool) string {
	if isToday(due) {
		return todayPriColor.Sprintf("today")
	} else if isTomorrow(due) {
		return tomorrowPriColor.Sprintf("tomorrow")
	} else if isPastDue(due) && !completed {
		return overduePriColor.Sprintf(due.Format("Mon Jan 02"))
	}
	return otherDuePriColor.Sprintf(due.Format("Mon Jan 02"))
}

func (f *ScreenPrinter) formatSubject(subject string, isPriority bool) string {
	splitted := strings.Split(subject, " ")

	if isPriority {
		return f.printPrioritySubject(splitted)
	}
	return f.printSubject(splitted)
}

func (f *ScreenPrinter) printPrioritySubject(splitted []string) string {
	coloredWords := []string{}
	for _, word := range splitted {
		if projectRegex.MatchString(word) {
			coloredWords = append(coloredWords, taskTextProjectPriWordColor.Sprint(word))
		} else if tagRegex.MatchString(word) {
			coloredWords = append(coloredWords, taskTextTagWordPriColor.Sprint(word))
		} else {
			coloredWords = append(coloredWords, taskTextPriColor.Sprint(word))
		}
	}
	return strings.Join(coloredWords, " ")
}

func (f *ScreenPrinter) printSubject(splitted []string) string {
	coloredWords := []string{}
	for _, word := range splitted {
		if projectRegex.MatchString(word) {
			coloredWords = append(coloredWords, taskTextProjectWordColor.Sprint(word))
		} else if tagRegex.MatchString(word) {
			coloredWords = append(coloredWords, taskTextTagWordColor.Sprint(word))
		} else {
			coloredWords = append(coloredWords, taskTextColor.Sprint(word))
		}
	}
	return strings.Join(coloredWords, " ")
}
