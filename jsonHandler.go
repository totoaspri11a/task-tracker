package main

import (
	"fmt"
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	fmt.Println("Error checking if file exists.", err)

	return false
}

func createJsonFile(filename string) {
	// create a json file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating a json file.", err)
		return
	}
	defer file.Close()

	// Escribe un array JSON vacío
	_, err = file.WriteString("[]")
	if err != nil {
		fmt.Println("Error writing to JSON file:", err)
		return
	}

	fmt.Println("Created JSON file:", file.Name())
}
