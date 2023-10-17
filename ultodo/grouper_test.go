package ultodo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGroupByTag(t *testing.T) {
	assert := assert.New(t)

	list := SetUpTestMemoryTodoList()
	grouper := &Grouper{}
	grouped := grouper.GroupByTag(list.Todos())

	assert.Equal(2, len(grouped.Groups["root"]), "")
	assert.Equal(1, len(grouped.Groups["more"]), "")
}

func TestGroupByProject(t *testing.T) {
	assert := assert.New(t)

	list := SetUpTestMemoryTodoList()
	grouper := &Grouper{}
	grouped := grouper.GroupByProject(list.Todos())

	assert.Equal(2, len(grouped.Groups["test1"]), "")
}

func TestGroupByTagWithPriorityFirst(t *testing.T) {
	assert := assert.New(t)

	var list []*Todo
	list = append(list, &Todo{Subject: "a - one", IsPriority: false})
	list = append(list, &Todo{Subject: "b - two", IsPriority: true})

	grouper := &Grouper{}
	grouped := grouper.GroupByTag(list)

	assert.Equal("b - two", grouped.Groups["No tags"][0].Subject)
}

func TestGroupByTagSortedByDueDate(t *testing.T) {
	assert := assert.New(t)

	var list []*Todo
	list = append(list, &Todo{Subject: "a - one", IsPriority: false, Due: time.Now().Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "b - two", IsPriority: false, Due: time.Now().AddDate(0, 0, -1).Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "c - three", IsPriority: false, Due: ""})

	grouper := &Grouper{}
	grouped := grouper.GroupByTag(list)

	assert.Equal("b - two", grouped.Groups["No tags"][0].Subject)
}

func TestGroupByTagSortedByDueDateWithNoDuePriority(t *testing.T) {
	assert := assert.New(t)

	var list []*Todo
	list = append(list, &Todo{Subject: "a - one", IsPriority: false, Due: time.Now().Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "b - two", IsPriority: false, Due: time.Now().AddDate(0, 0, -1).Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "c - three", IsPriority: true, Due: ""})

	grouper := &Grouper{}
	grouped := grouper.GroupByTag(list)

	assert.Equal("c - three", grouped.Groups["No tags"][0].Subject)
}

func TestGroupByTagSortedByDueDateWithPriority(t *testing.T) {
	assert := assert.New(t)

	var list []*Todo
	list = append(list, &Todo{Subject: "a - one", IsPriority: true, Due: time.Now().Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "b - two", IsPriority: false, Due: time.Now().AddDate(0, 0, -1).Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "c - three", IsPriority: false, Due: ""})

	grouper := &Grouper{}
	grouped := grouper.GroupByTag(list)

	assert.Equal("a - one", grouped.Groups["No tags"][0].Subject)
}

func TestGroupByTagSortedByDueDateWithArchived(t *testing.T) {
	assert := assert.New(t)

	var list []*Todo
	list = append(list, &Todo{Subject: "a - one", IsPriority: true, Archived: true, Due: time.Now().Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "b - two", IsPriority: false, Archived: true, Due: time.Now().AddDate(0, 0, -1).Format(DATE_FORMAT)})
	list = append(list, &Todo{Subject: "c - three", IsPriority: false, Due: ""})

	grouper := &Grouper{}
	grouped := grouper.GroupByTag(list)

	assert.Equal("c - three", grouped.Groups["No tags"][0].Subject)
}
