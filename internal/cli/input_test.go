package cli

import (
	"io"
	"os"
	"testing"
)

func TestGetInput(t *testing.T) {
	// Test case 1: Input from arguments
	args := []string{"hello world"}
	reader := GetInput(args)
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("Failed to read from reader: %v", err)
	}
	if string(output) != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", string(output))
	}

	// Test case 2: Input from stdin (mocking stdin is tricky in parallel, so we might skip or use a pipe)
	// For simplicity in this unit test, we'll trust that os.Stdin is returned when args is empty.
	// We can verify the type/value if needed, but integration tests usually handle the pipe case better.
	// However, we can check if it returns os.Stdin
	readerStdin := GetInput([]string{})
	if readerStdin != os.Stdin {
		t.Error("Expected GetInput to return os.Stdin when no args provided")
	}
}
