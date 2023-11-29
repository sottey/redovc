package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"projects/redo.vc/web/models"
)

const (
	filePath = "./tasks.json"
)

func ReadTasks() ([]models.Task, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var tasks []models.Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tasks, nil
}

func WriteTasks(tasks []models.Task) error {
	fileData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = os.WriteFile(filePath, fileData, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
