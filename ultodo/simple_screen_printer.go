package ultodo

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
type SimpleScreenPrinter struct {
	Writer         *io.Writer
	UnicodeSupport bool
}

// NewScreenPrinter creates a new screeen printer.
func NewSimpleScreenPrinter(unicodeSupport bool) *SimpleScreenPrinter {
	w := new(io.Writer)
	formatter := &SimpleScreenPrinter{Writer: w, UnicodeSupport: unicodeSupport}
	return formatter
}

// Print prints the output of ultodo to the terminal screen.
func (f *SimpleScreenPrinter) Print(groupedTodos *GroupedTodos, printNotes bool, showStatus bool) {
	var keys []string
	for key := range groupedTodos.Groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	tabby := tabby.NewCustom(tabwriter.NewWriter(color.Output, 0, 0, 2, ' ', 0))
	tabby.AddLine()
	for _, key := range keys {
		tabby.AddLine(key)
		for _, todo := range groupedTodos.Groups[key] {
			f.printTodo(tabby, todo, printNotes, showStatus)
		}
		tabby.AddLine()
	}
	tabby.Print()
}

func (f *SimpleScreenPrinter) printTodo(tabby *tabby.Tabby, todo *Todo, printNotes bool, showStatus bool) {
	if showStatus {
		tabby.AddLine(
			f.formatID(todo.ID, todo.IsPriority),
			f.formatCompleted(todo.Completed),
			f.formatInformation(todo),
			f.formatDue(todo.Due, todo.IsPriority, todo.Completed),
			f.formatSubject(todo.Subject, todo.IsPriority))
	} else {
		tabby.AddLine(
			f.formatID(todo.ID, todo.IsPriority),
			f.formatCompleted(todo.Completed),
			f.formatDue(todo.Due, todo.IsPriority, todo.Completed),
			f.formatSubject(todo.Subject, todo.IsPriority))
	}
	if printNotes {
		for nid, note := range todo.Notes {
			tabby.AddLine("  ", strconv.Itoa(nid), note)
		}
	}
}

func (f *SimpleScreenPrinter) formatID(ID int, isPriority bool) string {
	if isPriority {
		return strconv.Itoa(ID)
	}
	return strconv.Itoa(ID)
}

func (f *SimpleScreenPrinter) formatCompleted(completed bool) string {
	if completed {
		if f.UnicodeSupport {
			return "[✔]"
		} else {
			return "[x]"
		}
	}
	return "[ ]"
}

func (f *SimpleScreenPrinter) formatDue(due string, isPriority bool, completed bool) string {
	if due == "" {
		return "          "
	}
	dueTime, _ := time.Parse(DATE_FORMAT, due)

	if isPriority {
		return f.printPriorityDue(dueTime, completed)
	}
	return f.printDue(dueTime, completed)
}

func (f *SimpleScreenPrinter) formatInformation(todo *Todo) string {
	var information []string
	if todo.IsPriority {
		information = append(information, "*")
	} else {
		information = append(information, " ")
	}
	if todo.HasNotes() {
		information = append(information, "N")
	} else {
		information = append(information, " ")
	}
	if todo.Archived {
		information = append(information, "A")
	} else {
		information = append(information, " ")
	}
	return strings.Join(information, "")
}

func (f *SimpleScreenPrinter) printDue(due time.Time, completed bool) string {
	if isToday(due) {
		return "today     "
	} else if isTomorrow(due) {
		return "tomorrow  "
	} else if isPastDue(due) && !completed {
		return due.Format("Mon Jan 02")
	}
	return due.Format("Mon Jan 02")
}

func (f *SimpleScreenPrinter) printPriorityDue(due time.Time, completed bool) string {
	if isToday(due) {
		return "today     "
	} else if isTomorrow(due) {
		return "tomorrow  "
	} else if isPastDue(due) && !completed {
		return due.Format("Mon Jan 02")
	}
	return due.Format("Mon Jan 02")
}

func (f *SimpleScreenPrinter) formatSubject(subject string, isPriority bool) string {
	splitted := strings.Split(subject, " ")

	if isPriority {
		return f.printPrioritySubject(splitted)
	}
	return f.printSubject(splitted)
}

func (f *SimpleScreenPrinter) printPrioritySubject(splitted []string) string {
	coloredWords := []string{}
	for _, word := range splitted {
		if projectRegex.MatchString(word) {
			coloredWords = append(coloredWords, word)
		} else if tagRegex.MatchString(word) {
			coloredWords = append(coloredWords, word)
		} else {
			coloredWords = append(coloredWords, word)
		}
	}
	return strings.Join(coloredWords, " ")
}

func (f *SimpleScreenPrinter) printSubject(splitted []string) string {
	coloredWords := []string{}
	for _, word := range splitted {
		if projectRegex.MatchString(word) {
			coloredWords = append(coloredWords, word)
		} else if tagRegex.MatchString(word) {
			coloredWords = append(coloredWords, word)
		} else {
			coloredWords = append(coloredWords, word)
		}
	}
	return strings.Join(coloredWords, " ")
}
