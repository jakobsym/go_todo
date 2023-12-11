package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const TodoFileName = "test.json" // TODO: Hardcode a file name (Change later)

func main() {
	/* Can use flag.Usage at top to create custom -h message */
	//var newTask = flag.String("task", "", "tetse")

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
		fmt.Println("task = ", *task)
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
