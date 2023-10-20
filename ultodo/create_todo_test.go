package ultodo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodoNoDue(t *testing.T) {
	assert := assert.New(t)

	parser := &InputParser{}
	filter, _ := parser.Parse("subject with a +project and #tag")

	todo, _ := CreateTodo(filter)
	assert.Equal("subject with a +project and #tag", todo.Subject)
	assert.Equal("", todo.Due)
	assert.Equal(false, todo.IsPriority)
	assert.Equal(false, todo.Archived)
	assert.Equal(false, todo.Completed)
	assert.Equal([]string{"project"}, todo.Projects)
	assert.Equal([]string{"tag"}, todo.Tags)
	assert.Equal("", todo.Status)
}

func TestCreateTodoWithDue(t *testing.T) {
	assert := assert.New(t)
	tomorrow := time.Now().AddDate(0, 0, 1)
	tomorrowString := tomorrow.Format("Jan2")

	parser := &InputParser{}
	filter, _ := parser.Parse("subject with a +project and #tag due:" + tomorrowString)
	fmt.Println(filter)
	todo, _ := CreateTodo(filter)
	assert.Equal("subject with a +project and #tag", todo.Subject)
	fmt.Println("todo.due: " + todo.Due)
	fmt.Println("tomorrow.Format(DATE_FORMAT): " + tomorrow.Format(DATE_FORMAT))
	assert.Equal(tomorrow.Format(DATE_FORMAT), todo.Due)
	assert.Equal(false, todo.IsPriority)
	assert.Equal(false, todo.Archived)
	assert.Equal(false, todo.Completed)
	assert.Equal([]string{"project"}, todo.Projects)
	assert.Equal([]string{"tag"}, todo.Tags)
	assert.Equal("", todo.Status)
}

func TestCreateTodoWithStatusAndPriority(t *testing.T) {
	assert := assert.New(t)

	parser := &InputParser{}
	filter, _ := parser.Parse("subject with a +project and #tag status:waiting priority:true")

	todo, _ := CreateTodo(filter)
	assert.Equal("subject with a +project and #tag", todo.Subject)
	assert.Equal("", todo.Due)
	assert.Equal(true, todo.IsPriority)
	assert.Equal(false, todo.Archived)
	assert.Equal(false, todo.Completed)
	assert.Equal([]string{"project"}, todo.Projects)
	assert.Equal([]string{"tag"}, todo.Tags)
	assert.Equal("waiting", todo.Status)
}

func TestCreateCompletedTodo(t *testing.T) {
	assert := assert.New(t)

	parser := &InputParser{}
	filter, _ := parser.Parse("lunch with bob completed:true")

	todo, _ := CreateTodo(filter)
	assert.Equal("lunch with bob", todo.Subject)
	assert.Equal("", todo.Due)
	assert.Equal(false, todo.IsPriority)
	assert.Equal(false, todo.Archived)
	assert.Equal(true, todo.Completed)
	assert.Equal([]string{}, todo.Projects)
	assert.Equal([]string{}, todo.Tags)
	assert.Equal("completed", todo.Status)
}
