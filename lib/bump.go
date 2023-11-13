package lib

import (
	"time"
)

// SCOTODO: Fix this
func (a *App) Bump() {
	a.load()

	today := time.Now().Format(DATE_FORMAT)

	for _, todo := range a.TodoList.Todos() {
		due, _ := time.Parse(DATE_FORMAT, todo.Due)
		if todo.Due != "" && isPastDue(due) {
			todo.Due = today
		}
	}
	a.save()
}
