package ultodo

// Store is the interface for ultodo todos.
type Store interface {
	GetLocation() string
	LocalTodosFileExists() bool
	Initialize()
	Load() ([]*Todo, error)
	Save(todos []*Todo)
}
