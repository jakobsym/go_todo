package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

func main() {
	/* Can use flag.Usage at top to create custom -h message */

	var TodoFileName = ".todo.json"
	// create via -> < export TODO_FILENAME=new-todo.json >  from your cli
	if os.Getenv("TODO_FILENAME") != "" {
		TodoFileName = os.Getenv("TODO_FILENAME")
	}

	var task = flag.String("task", "", "Task included in Todo list")
	var list = flag.Bool("list", false, "List todo list task(s)")
	var complete = flag.Int("complete", 0, "Task to be completed")

	flag.Parse()

	l := &todo.List{} // & because (l *List) type

	if err := l.Get(TodoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	// Display all todo items in List
	case *list:
		// List to do items
		fmt.Print(l)

	// Mark Complete
	case *complete > 0:
		// Complete item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save changes
		if err := l.Save(TodoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add Task
		l.Add(*task)

		// Save Changes
		if err := l.Save(TodoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}
}
