package cli

import (
	"io"
	"os"
	"strings"
)

// GetInput returns a reader for the input data.
// If arguments are provided, it treats the first argument as the input string.
// Otherwise, it reads from stdin.
func GetInput(args []string) io.Reader {
	if len(args) > 0 {
		return strings.NewReader(args[0])
	}
	return os.Stdin
}
