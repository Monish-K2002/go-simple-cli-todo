package main

import (
	"fmt"
)

func main() {
	todos := todos{}
	err := LoadFile(&todos)
	if err != nil {
		fmt.Println(err)
	}
	parse := Parse()
	parse.Execute(&todos)
	err = AppendFile(&todos)
	if err != nil {
		fmt.Println(err)
	}
}
