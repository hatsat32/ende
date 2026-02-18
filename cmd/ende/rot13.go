package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hatsat32/ende/internal/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(enrot13Cmd)
	rootCmd.AddCommand(derot13Cmd)
}

func rot13(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return 'a' + (r-'a'+13)%26
	}
	if r >= 'A' && r <= 'Z' {
		return 'A' + (r-'A'+13)%26
	}
	return r
}

func runRot13(args []string) error {
	input := cli.GetInput(args)
	reader := bufio.NewReader(input)
	
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(string(rot13(r)))
	}
	return nil
}

var enrot13Cmd = &cobra.Command{
	Use:   "enrot13 [input]",
	Short: "Encode data to ROT13",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRot13(args)
	},
}

var derot13Cmd = &cobra.Command{
	Use:   "derot13 [input]",
	Short: "Decode data from ROT13",
	RunE: func(cmd *cobra.Command, args []string) error {
		// ROT13 is symmetric
		return runRot13(args)
	},
}
