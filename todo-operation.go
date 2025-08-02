package main

import (
	//"os"
	"fmt"
	"os"
	"slices"
	"time"
	//"fmt"
	"text/tabwriter"
)

type todo struct {
	S_No        int
	Title       string
	Status      string
	CreatedAt   time.Time
	CompletedAt interface{}
}

type todos []todo

func (t *todos) add(title string) {
	fmt.Println("title", title)

	newTodo := todo{
		S_No:        len(*t) + 1,
		Title:       title,
		Status:      "pending",
		CreatedAt:   time.Now(),
		CompletedAt: "",
	}
	fmt.Println("New todo", newTodo)
	//Check if the given Title already exists and if exists replace the index
	existingTitle := validateTitleIndex(title, *t)
	if existingTitle != -1 {
		(*t)[existingTitle] = newTodo
		return
	}

	*t = append(*t, newTodo)
}

func (t *todos) updateStatus(title string) {
	existingTitle := validateTitleIndex(title, *t)
	if existingTitle == -1 {
		return
	}

	status := (*t)[existingTitle].Status
	if status == "pending" {
		(*t)[existingTitle].CompletedAt = time.Now()
		(*t)[existingTitle].Status = "completed"
	} else {
		(*t)[existingTitle].CompletedAt = ""
		(*t)[existingTitle].Status = "pending"
	}

}

func (t *todos) deleteIndex(index int) {
	existingIndex := validateIndex(index, *t)
	if existingIndex == -1 {
		fmt.Println("Index Not found")
		return
	}

	*t = slices.Delete(*t, existingIndex, existingIndex+1)
}

func (t *todos) list() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Print("\n")
	fmt.Fprintln(w, "Title\tStatus\tCreatedAt\tCompletedAt")
	for _, t := range *t {
		fmt.Fprintf(w, "%v\t%s\t%s\t%s\t%s\n", t.S_No, t.Title, t.Status, t.CreatedAt, t.CompletedAt)
	}
	err := w.Flush()
	fmt.Print("\n")
	if err != nil {
		panic(err)
	}
}

func validateIndex(index int, data todos) int {
	return slices.IndexFunc(data, func(t todo) bool { return t.S_No == index })
}

func validateTitleIndex(title string, data todos) int {
	return slices.IndexFunc(data, func(n todo) bool {
		return n.Title == title
	})
}
