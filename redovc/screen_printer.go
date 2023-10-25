package redovc

import (
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

// NewScreenPrinter creates a new screeen printer.
func NewScreenPrinter(unicodeSupport bool) *ScreenPrinter {
	w := new(io.Writer)
	formatter := &ScreenPrinter{Writer: w, UnicodeSupport: unicodeSupport}
	return formatter
}

// Print prints the output of redovc to the terminal screen.
func (f *ScreenPrinter) Print(groupedTodos *GroupedTodos, printNotes bool, showInfoFlags bool) {
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
			f.printTodo(tabby, todo, printNotes, showInfoFlags)
		}
		tabby.AddLine()
	}
	tabby.Print()
}

func (f *ScreenPrinter) printTodo(tabby *tabby.Tabby, todo *Todo, printNotes bool, showInfoFlags bool) {
	if showInfoFlags {
		tabby.AddLine(
			f.formatID(todo.ID, todo.IsPriority),
			f.formatCompleted(todo.Completed),
			f.formatInfoFlags(todo),
			f.formatDue(todo.Due, todo.IsPriority, todo.Completed),
			f.formatStatus(todo.Status, todo.IsPriority),
			f.formatSubject(todo.Subject, todo.IsPriority))
	} else {
		tabby.AddLine(
			f.formatID(todo.ID, todo.IsPriority),
			f.formatCompleted(todo.Completed),
			f.formatDue(todo.Due, todo.IsPriority, todo.Completed),
			f.formatStatus(todo.Status, todo.IsPriority),
			f.formatSubject(todo.Subject, todo.IsPriority))
	}

	if printNotes {
		for nid, note := range todo.Notes {
			tabby.AddLine(
				"  "+noteIDColor.Sprint(strconv.Itoa(nid)),
				noteTextColor.Sprint(""),
				noteTextColor.Sprint(""),
				noteTextColor.Sprint(""),
				noteTextColor.Sprint(""),
				noteTextColor.Sprint(note))
		}
	}
}

func (f *ScreenPrinter) formatID(ID int, isPriority bool) string {
	if isPriority {
		return taskIDPriColor.Sprint(strconv.Itoa(ID))
	}
	return taskIDColor.Sprint(strconv.Itoa(ID))
}

func (f *ScreenPrinter) formatCompleted(completed bool) string {
	if completed {
		if f.UnicodeSupport {
			return completedColor.Sprint("[✔]")
		}
		return completedColor.Sprint("[x]")
	}
	return completedColor.Sprint("[ ]")
}

func (f *ScreenPrinter) formatDue(due string, isPriority bool, completed bool) string {
	if due == "" {
		return whiteFg.Sprint("          ")
	}
	dueTime, _ := time.Parse(DATE_FORMAT, due)

	if isPriority {
		return f.printPriorityDue(dueTime, completed)
	}
	return f.printDue(dueTime, completed)
}

func (f *ScreenPrinter) formatStatus(status string, isPriority bool) string {
	if status == "" {
		return statusColor.Sprint("          ")
	}

	if len(status) < 10 {
		for x := len(status); x <= 10; x++ {
			status += " "
		}
	}

	statusRune := []rune(status)

	if isPriority {
		return statusPriColor.Sprintf("%-10v", string(statusRune[0:10]))
	}
	return statusColor.Sprintf("%-10s", string(statusRune[0:10]))
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
		return todayColor.Sprint("today     ")
	} else if isTomorrow(due) {
		return tomorrowColor.Sprint("tomorrow  ")
	} else if isPastDue(due) && !completed {
		return overdueColor.Sprint(due.Format("Mon Jan 02"))
	}

	return otherDue.Sprint(due.Format("Mon Jan 02"))
}

func (f *ScreenPrinter) printPriorityDue(due time.Time, completed bool) string {
	if isToday(due) {
		return todayPriColor.Sprint("today     ")
	} else if isTomorrow(due) {
		return tomorrowPriColor.Sprint("tomorrow  ")
	} else if isPastDue(due) && !completed {
		return overduePriColor.Sprint(due.Format("Mon Jan 02"))
	}
	return otherDuePriColor.Sprint(due.Format("Mon Jan 02"))
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
