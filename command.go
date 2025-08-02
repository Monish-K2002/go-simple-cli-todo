package main

import (
	"flag"
	"fmt"
)

type Args struct {
	Add    string
	Update string
	Delete int
	List   bool
}

func Parse() *Args {
	args := Args{}
	flag.StringVar(&args.Add, "Add", "", "add string")
	flag.StringVar(&args.Update, "Update", "", "Update the status of existing todo list")
	flag.IntVar(&args.Delete, "Delete", 0, "Delete the Todo Entry based on Index")
	flag.BoolVar(&args.List, "List", false, "List the values from JSON")
	flag.Parse()

	return &args
}

func (args *Args) Execute(t *todos) {
	todos := t
	switch {
	case args.Add != "":
		todos.add(args.Add)
		break
	case args.Update != "":
		todos.updateStatus(args.Update)
		break
	case args.Delete != 0:
		todos.deleteIndex(args.Delete)
	case args.List:
		todos.list()
		break
	default:
		fmt.Println("Error: You must specify -Add, -Update, -List: ")
	}
}
