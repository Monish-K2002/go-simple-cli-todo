package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func LoadFile(data *todos) error {
	fileName := "todo.json"
	checkFileValidity(fileName)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}(file)

	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(fileContent, data)

	return err
}

func AppendFile(data *todos) error {
	fileName := "todo.json"
	fileContent, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileName, fileContent, 0644)
	return err
}

func checkFileValidity(fileName string) {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Creating the file %s \n", fileName)
			return
		}
	} else {
		log.Printf("Opening the file %s \n", fileName)
		return
	}
}
