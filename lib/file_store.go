package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var useVikunja = false

const TodosJSONFile = ".redovc.todos.json"
const VikunjaBaseURL = "http://192.168.7.247:8084/api/v1"
const VikunjaToken = "tk_ab722f23d9ad9377e6881a55348c672e4b47711b"

const TodosThemeFile = ".redovc.theme.json"

// FileStore is the main struct of this file.
type FileStore struct {
	Loaded bool
}

// NewFileStore is creating a new file store.
func NewFileStore() *FileStore {
	return &FileStore{Loaded: false}
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

// Initialize is initializing a new .redovc.todos.json file.
func (f *FileStore) Initialize() {
	if useVikunja {
		fmt.Println("When using Vikunja, initialization is not needed")
	} else {
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
	var todos []*Todo
	var getVikError error

	if useVikunja {
		todos, getVikError = LoadVikunja()
		if getVikError != nil {
			return nil, getVikError
		}
	} else {
		data, err := os.ReadFile(f.GetLocation())
		if err != nil {
			fmt.Println("No todo file found!")
			fmt.Println("Initialize a new todo repo by running 'redovc init'")
			os.Exit(1)
			return nil, err
		}

		jerr := json.Unmarshal(data, &todos)
		if jerr != nil {
			os.Exit(1)
			return nil, jerr
		}

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

	if useVikunja {
		SaveVikunja(todos)
	} else {
		data, _ := json.MarshalIndent(todos, "", "    ")

		if err := os.WriteFile(f.GetLocation(), []byte(data), 0644); err != nil {
			fmt.Println("Error Saving file: ", err)
		}
	}
}

func LoadVikunja() ([]*Todo, error) {

	// Call API and get data
	apiURL := VikunjaBaseURL + "/tasks/all"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+VikunjaToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Non-OK HTTP status: %d\n%s", resp.StatusCode, body)
		return nil, err
	}

	var vTasks []VikunjaTask
	err = json.NewDecoder(resp.Body).Decode(&vTasks)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return nil, err
	} else {
		// remap data to redovc format
		ret, convertErr := VikunjaToRedo(vTasks)
		if convertErr != nil {
			return nil, convertErr
		}

		// return data
		return ret, nil
	}
}

func SaveVikunja(todos []*Todo) error {
	// Get the vikunja data

	// remap data to Vikunja format

	// Save the data to vikunja

	return nil
}

func VikunjaToRedo(vData []VikunjaTask) ([]*Todo, error) {
	var ret []*Todo
	//fmt.Println(vData)

	for _, vTask := range vData {
		var rTask Todo

		// fmt.Printf("vTask: \n%v\n\nrTask:\n%v", vTask, rTask)
		rTask.ID = vTask.ID
		// rTask.UUID = vTask.Identifier
		rTask.Subject = vTask.Title
		// rTask.Due = vTask.DueDate
		rTask.Completed = vTask.Done
		// rTask.CompletedDate = vTask.DoneAt
		// rTask.IsPriority = (vTask.Priority == 0)
		// rTask.Notes = append(rTask.Notes, vTask.Description)

		ret = append(ret, &rTask)
	}

	return ret, nil
}
