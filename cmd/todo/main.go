package main

import (
	"fmt"
	"os"
	"strings"
	"todo"
)

const TodoFileName = "test.json" // TODO: Hardcode a file name (Change later)

func main() {
	l := &todo.List{} // & because (l *List) type

	if err := l.Get(TodoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	// Display all todo items in List
	case len(os.Args) == 1:
		// List to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// Add
	default:
		item := strings.Join(os.Args[1:], " ") // os.Args[] is great but not flexiable; Better to use 'flags'

		// add item list
		l.Add(item)
		// save the list
		if err := l.Save(TodoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
