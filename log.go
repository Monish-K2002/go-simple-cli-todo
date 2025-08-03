package main

import (
	"log"
	"os"
)

func CreateLog() (*os.File, error) {
	fileName := "todo-log.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	return file, nil
}
