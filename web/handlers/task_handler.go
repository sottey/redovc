package handlers

import (
	"encoding/json"
	"net/http"

	"projects/redo.vc/web/models"
	"projects/redo.vc/web/utils"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := utils.ReadTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := utils.ReadTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate unique ID for the new task
	task.ID = 999 // SCOCHANGE: // utils.GenerateUniqueID()

	tasks = append(tasks, task)

	err = utils.WriteTasks(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	taskID := 999 // SCOCHANGE: // := vars["id"]

	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := utils.ReadTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			// Update the task properties
			tasks[i].Project = updatedTask.Project
			tasks[i].Tags = updatedTask.Tags
			//tasks[i].Due = updatedTask.Due
			tasks[i].Completed = updatedTask.Completed
			tasks[i].Archived = updatedTask.Archived
			tasks[i].Subject = updatedTask.Subject

			break
		}
	}

	err = utils.WriteTasks(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	taskID := 999 // SCOCHANGE: // vars["id"]

	tasks, err := utils.ReadTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			// Remove the task from the slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	err = utils.WriteTasks(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
