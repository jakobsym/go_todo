package todo_test

// to import 'package todo'
// just import the path to the todo.go file

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	i := 1
	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done != false {
		t.Errorf("Expected %v, got %v instead", false, l[0].Done)
	}

	l.Complete(i)

	if l[0].Done != true {
		t.Errorf("Expected %v, got %v instead", false, l[0].Done)
	}
}

func TestDelete(t *testing.T) {
	// create a list of items
	l := todo.List{}
	listSize := 2

	tasks := []string{
		"New Task 1", // 0
		"New Task 2",
		"New Task 3", // 2
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(listSize)

	if len(l) != listSize {
		t.Errorf("Expected %q, got %q instead", listSize, len(l))
	}

	// check that the taks at index 1, != task at index 2 (if deletion correct, these should ==)
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead", tasks[2], l[1].Task)
	}
}

// test both Save(), Get()
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "") // create temp. file

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name()) // close file at end of executing function `defer`

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	// load saved list l1 into l2, they should have same tasks
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l1[0].Task, l2[0].Task)
	}

}
