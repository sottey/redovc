package redovc

// Store is the interface for redovc todos.
type Store interface {
	GetLocation() string
	LocalTodosFileExists() bool
	Initialize()
	Load() ([]*Todo, error)
	Save(todos []*Todo)
}
