package utils

import "goto/interfaces"

func IsContain(name string, tasks []interfaces.Task) bool {
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		if task.Name == name {
			return true
		}
		for j := 0; j < len(task.Alias); j++ {
			if task.Alias[j] == name {
				return true
			}
		}
	}
	return false
}

func FindIndex(name string, tasks []interfaces.Task) int {
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		if task.Name == name {
			return i
		}
		for j := 0; j < len(task.Alias); j++ {
			if task.Alias[j] == name {
				return i
			}
		}
	}
	return -1
}
