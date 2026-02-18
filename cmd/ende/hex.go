package main

import (
	"encoding/hex"
	"io"
	"os"

	"github.com/hatsat32/ende/internal/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(enhexCmd)
	rootCmd.AddCommand(dehexCmd)
	rootCmd.AddCommand(enb16Cmd)
	rootCmd.AddCommand(deb16Cmd)
}

var enhexCmd = &cobra.Command{
	Use:   "enhex [input]",
	Short: "Encode data to Hexadecimal",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runHexEncoder(args)
	},
}

var dehexCmd = &cobra.Command{
	Use:   "dehex [input]",
	Short: "Decode data from Hexadecimal",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runHexDecoder(args)
	},
}

var enb16Cmd = &cobra.Command{
	Use:   "enb16 [input]",
	Short: "Encode data to Base16 (Hex alias)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runHexEncoder(args)
	},
}

var deb16Cmd = &cobra.Command{
	Use:   "deb16 [input]",
	Short: "Decode data from Base16 (Hex alias)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runHexDecoder(args)
	},
}

func runHexEncoder(args []string) error {
	input := cli.GetInput(args)
	encoder := hex.NewEncoder(os.Stdout)
	if _, err := io.Copy(encoder, input); err != nil {
		return err
	}
	return nil
}

func runHexDecoder(args []string) error {
	input := cli.GetInput(args)
	decoder := hex.NewDecoder(input)
	if _, err := io.Copy(os.Stdout, decoder); err != nil {
		return err
	}
	return nil
}
