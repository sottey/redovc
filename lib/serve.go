package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Serves todos via http
func (a *App) Serve(ipPort string) {
	http.HandleFunc("/", serveJSON)
	http.HandleFunc("/html", serveHTML)
	http.HandleFunc("/json", serveJSON)
	http.HandleFunc("/csv", serveCSV)

	err := http.ListenAndServe(ipPort, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
		os.Exit(0)
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {

	app := NewApp()
	list, _ := app.GetJSONFileContents()

	_, output := JSON2HtmlTable(string(list), nil, nil)
	io.WriteString(w, output)
}

func serveJSON(w http.ResponseWriter, r *http.Request) {

	app := NewApp()
	list, _ := app.GetJSONFileContents()
	io.WriteString(w, string(list))
}

// SCOTODO: Finish implementation
func serveCSV(w http.ResponseWriter, r *http.Request) {
	app := NewApp()
	list, _ := app.GetJSONFileContents()

	output := JSONtoCSV(list)
	io.WriteString(w, output)
}

/*
	tags
	recur
	id
	status
	archived
	is_priority
	recur_until
	subject
	completed
	notes
	uuid
	due
	completed_date
	prev_recur_todo_uuid
	projects
*/
