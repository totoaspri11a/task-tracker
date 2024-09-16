package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Definición de la estructura Task
type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// Carga las tareas desde un archivo JSON
func loadTasks(filename string) ([]Task, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Función para obtener el siguiente ID basado en el ID más alto en la lista
func getNextID(tasks []Task) string {
	maxID := 0

	for _, task := range tasks {
		id, err := strconv.Atoi(task.Id)
		if err != nil {
			fmt.Println("Error convirtiendo ID a entero:", err)
			continue
		}
		if id > maxID {
			maxID = id
		}
	}

	return fmt.Sprintf("%d", maxID+1)
}

func addTask(tasks []Task, description string) []Task {
	newTask := Task{
		Id:          getNextID(tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	return append(tasks, newTask)
}

// Función para actualizar una tarea
func updateTask(tasks []Task, id string, newValue string) []Task {
	// Definir los valores válidos para el estado
	validStatuses := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	for i, task := range tasks {
		if task.Id == id {
			if _, isValidStatus := validStatuses[newValue]; isValidStatus {
				// Si el nuevo valor es un estado válido, actualizar el estado
				task.Status = newValue
			} else {
				// De lo contrario, actualizar la descripción
				task.Description = newValue
			}
			task.UpdatedAt = time.Now().Format(time.RFC3339)
			tasks[i] = task // Actualizar la tarea en la lista
			break
		}
	}
	return tasks
}

// Función para eliminar una tarea
func deleteTask(tasks []Task, id string) []Task {
	for i, task := range tasks {
		if task.Id == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}
