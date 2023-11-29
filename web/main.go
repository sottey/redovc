package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"projects/redo.vc/web/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
