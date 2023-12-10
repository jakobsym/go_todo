package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const TodoFileName = "test.json" // TODO: Hardcode a file name (Change later)

func main() {
	task := flag.String("task", "", "Task included in Todo list")
	list := flag.Bool("list", false, "List todo list task(s)")
	complete := flag.Int("complete", 0, "Task to be completed")

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
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}

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
