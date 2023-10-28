package redovc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// TodosJSONFile is the filename to store todos in
const TodosJSONFile = ".redovc.todos.json"

// TodosThemeFile is the filename containing theme information
const TodosThemeFile = ".redovc.theme.json"

// FileStore is the main struct of this file.
type FileStore struct {
	Loaded bool
}

// NewFileStore is creating a new file store.
func NewFileStore() *FileStore {
	return &FileStore{Loaded: false}
}

// Initialize is initializing a new .redovc.todos.json file.
func (f *FileStore) Initialize() {
	if f.LocalTodosFileExists() {
		fmt.Println("It looks like a .redovc.todos.json file already exists!  Doing nothing.")
		os.Exit(0)
	}
	if err := os.WriteFile(TodosJSONFile, []byte("[]"), 0644); err != nil {
		fmt.Println("Error writing json file", err)
		os.Exit(1)
	}

	if f.LocalThemeFileExists() {
		fmt.Println("It looks like a .redovc.theme.json file already exists!  Doing nothing.")
		os.Exit(0)
	}

	fmt.Println("creating theme file")
	if err := os.WriteFile(TodosThemeFile, []byte(themeTemplate), 0644); err != nil {
		fmt.Println("Error writing theme file", err)
		os.Exit(1)
	}
}

// Returns if a local .redovc.todos.json file exists in the current dir.
func (f *FileStore) LocalTodosFileExists() bool {
	dir, _ := os.Getwd()
	localrepo := filepath.Join(dir, TodosJSONFile)
	_, err := os.Stat(localrepo)
	return err == nil
}

// Returns if a local .redovc.theme.json file exists in the current dir.
func (f *FileStore) LocalThemeFileExists() bool {
	dir, _ := os.Getwd()
	localrepo := filepath.Join(dir, TodosThemeFile)
	_, err := os.Stat(localrepo)
	return err == nil
}

// Load is loading a .redovc.todos.json file, either from cwd, or the home directory
func (f *FileStore) Load() ([]*Todo, error) {
	data, err := os.ReadFile(f.GetLocation())
	if err != nil {
		fmt.Println("No todo file found!")
		fmt.Println("Initialize a new todo repo by running 'redovc init'")
		os.Exit(1)
		return nil, err
	}

	var todos []*Todo
	jerr := json.Unmarshal(data, &todos)
	if jerr != nil {
		fmt.Println("Error reading json data", jerr)
		os.Exit(1)
		return nil, jerr
	}
	f.Loaded = true

	return todos, nil
}

// Save is saving a .redovc.todos.json file.
func (f *FileStore) Save(todos []*Todo) {
	// ensure UUID is set for todos at save time
	for _, todo := range todos {
		if todo.UUID == "" {
			todo.UUID = newUUID()
		}
	}

	data, _ := json.MarshalIndent(todos, "", "    ")

	if err := os.WriteFile(f.GetLocation(), []byte(data), 0644); err != nil {
		fmt.Println("Error writing json file", err)
	}
}

// GetLocation is returning the location of the .redovc.todos.json file.
func (f *FileStore) GetLocation() string {
	if f.LocalTodosFileExists() {
		dir, _ := os.Getwd()
		localrepo := filepath.Join(dir, TodosJSONFile)
		return localrepo
	}
	return fmt.Sprintf("%s/%s", UserDataDir(), TodosJSONFile)
}
