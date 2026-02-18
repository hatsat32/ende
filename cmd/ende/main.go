package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ende",
	Short: "ende is a CLI tool for encoding and decoding data streams",
	Long: `ende supports various encoding formats such as Base64, Base32, Hex, HTML Entities, URL encoding, and ROT13.
It reads from standard input or command line arguments and writes to standard output.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
