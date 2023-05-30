package utils

import (
	"encoding/json"
	"goto/interfaces"
	"io"
	"os"
)

func ReadFile(file string) ([]interfaces.Task, error) {
	tasks := []interfaces.Task{}
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return tasks, err
	}
	defer f.Close()

	dataStr, err := io.ReadAll(f)
	if err != nil {
		return tasks, err
	}
	if string(dataStr) == "" {
		return tasks, nil
	}

	err = json.Unmarshal(dataStr, &tasks)
	return tasks, err
}

func WriteFile(file string, tasks []interfaces.Task) error {
	dataStr, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, dataStr, 0644)
	return err
}
