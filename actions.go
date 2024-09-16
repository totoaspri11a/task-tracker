package main

import "fmt"

func processActions(action string, args []string, tasks []Task) ([]Task, error) {
	switch action {
	case "add":
		if len(args) < 1 {
			return tasks, fmt.Errorf("uso: add <descripción>")
		}
		description := args[0]
		tasks = addTask(tasks, description)
		fmt.Println("Tarea agregada con éxito.")

	case "update":
		if len(args) < 2 {
			return tasks, fmt.Errorf("uso: update <id> <nuevo_estado>")
		}
		id := args[0]
		newValue := args[1]
		tasks = updateTask(tasks, id, newValue)
		fmt.Println("Tarea actualizada con éxito.")

	case "delete":
		if len(args) < 1 {
			return tasks, fmt.Errorf("uso: delete <id>")
		}
		id := args[0]
		tasks = deleteTask(tasks, id)
		fmt.Println("Tarea eliminada con éxito.")

	default:
		return tasks, fmt.Errorf("acción no válida. uso: <add|update|delete> [args]")
	}

	return tasks, nil
}
