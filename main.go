package main

import (
	"log"
)

func main() {
	logFile, err := CreateLog()
	defer logFile.Close()
	log.Printf("Beginning the operation")
	todos := todos{}
	err = LoadFile(&todos)
	if err != nil {
		log.Printf("Error loading todos file: %v", err)
	}
	parse := Parse()
	parse.Execute(&todos)
	err = AppendFile(&todos)
	if err != nil {
		log.Printf("Error while appending file: %v", err)
	}
	log.Printf("Done")
}
