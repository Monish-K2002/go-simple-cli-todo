package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadFile(data *todos) error {
	fileName := "todo.json"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

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
