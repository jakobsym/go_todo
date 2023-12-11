package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// Integration Test Goals:
// - go build; compile into bin. file
// - execute bin. file with diff args and assert correct behavior

var (
	binName  = "todo"
	fileName = "testing_list.json"
)

// Calls 'go build' tool
// builds executable bin for tool
// execute using m.Run()
func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += " .exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

// Creates test cases
// uses subsets to execute tests
func TestTodoCLI(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd()
	//fmt.Println("dir = ", dir)

	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName) // path to build tool compiled

	t.Run("AddNewTaskFromArgs", func(t *testing.T) {
		os.Setenv("TODO_FILENAME", fileName)
		cmd := exec.Command(cmdPath, "-add", task) // execute compiled binary
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
	task2 := "test task number 2"

	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		os.Setenv("TODO_FILENAME", fileName)
		cmd := exec.Command(cmdPath, "-add")
		cmdStdin, err := cmd.StdinPipe() // create stdin pipe connection

		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdin, task2) // write to stdin pipe
		cmdStdin.Close()                // close stdin

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		os.Setenv("TODO_FILENAME", fileName)
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf(" 1: %s\n 2: %s\n", task, task2)

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}

	})

}
