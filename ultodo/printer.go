package ultodo

import "regexp"

var (
	projectRegex, _ = regexp.Compile(`\+[\p{L}\d_]+`)
	tagRegex, _     = regexp.Compile(`\#[\p{L}\d_]+`)
)

// Printer is an interface for printing grouped todos.
type Printer interface {
	Print(*GroupedTodos, bool, bool)
}
