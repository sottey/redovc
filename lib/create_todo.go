package lib

// CreateTodo will create a TodoItem from a Filter
func CreateTodo(filter *Filter) (*Todo, error) {
	todoItem := &Todo{
		Subject:    filter.Subject,
		Archived:   filter.Archived,
		IsPriority: filter.IsPriority,
		Completed:  filter.Completed,
		Projects:   filter.Projects,
		Tags:       filter.Tags,
		Due:        filter.Due,
		Status:     filter.LastStatus(),
		Recur:      filter.Recur,
		RecurUntil: filter.RecurUntil,
	}
	if todoItem.Completed {
		todoItem.Complete()
	}

	return todoItem, nil
}
