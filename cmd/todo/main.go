package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

func main() {
	/* Can use flag.Usage at top to create custom -h message */

	var TodoFileName = "todo_list.json"
	// create via -> < export TODO_FILENAME=new-todo.json >  from your cli
	if os.Getenv("TODO_FILENAME") != "" {
		TodoFileName = os.Getenv("TODO_FILENAME")
	}

	var addTask = flag.Bool("add", false, "Add task to todo list")
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
	case *addTask:
		// Add Task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
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

// returns string or potential error
func getTask(r io.Reader, args ...string) (string, error) {
	//fmt.Println("args = ", args)

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	//fmt.Println("s.Text() = ", s.Text())

	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank.\n")
	}

	return s.Text(), nil
}
