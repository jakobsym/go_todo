package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item // 1 indexed

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l

	// check out of todo list bounds
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist\n", i)
	}

	// mark as complete
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l

	// check out of todo list bounds
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist\n", i)
	}

	// ls = [1,2,3,4], i = 2 deleting '3'
	// below creates: [1,2,4]
	*l = append(ls[:i-1], ls[i:]...) // `...` := variadic function (can take any number of arguments)

	return nil
}

// Save list as JSON to a file
func (l *List) Save(fileName string) error {
	js, err := json.Marshal(l) // Converts to a JSON byte slice ([]byte)

	// err === nil if no error occurs

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, js, 0644)
}

// Retrieve JSON from file, convert to a `List`
func (l *List) Get(fileName string) error {
	file, err := os.ReadFile(fileName)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l) // converting JSON into Go struct
}
