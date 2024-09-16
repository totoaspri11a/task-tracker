package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: <command> [args]")
		return
	}

	// Asegúrate de que el primer argumento es un comando
	action := os.Args[1]
	args := os.Args[2:]
	filename := "db.json"
	var tasks []Task
	var err error

	if fileExists(filename) {
		tasks, err = loadTasks(filename)
		if err != nil {
			fmt.Println("Error cargando tareas:", err)
			return
		}
	} else {
		fmt.Println("Creando un archivo JSON.")
		createJsonFile(filename)
	}
	tasks, err = processActions(action, args, tasks)
	if err != nil {
		fmt.Println("Error procesando comandos:", err)
		return
	}

	err = saveTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error guardando tareas:", err)
		return
	}
}
