package main

import (
	"encoding/base32"
	"io"
	"os"

	"ende/internal/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(enb32Cmd)
	rootCmd.AddCommand(deb32Cmd)
}

var enb32Cmd = &cobra.Command{
	Use:   "enb32 [input]",
	Short: "Encode data to Base32",
	RunE: func(cmd *cobra.Command, args []string) error {
		input := cli.GetInput(args)
		encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
		if _, err := io.Copy(encoder, input); err != nil {
			return err
		}
		return encoder.Close()
	},
}

var deb32Cmd = &cobra.Command{
	Use:   "deb32 [input]",
	Short: "Decode data from Base32",
	RunE: func(cmd *cobra.Command, args []string) error {
		input := cli.GetInput(args)
		decoder := base32.NewDecoder(base32.StdEncoding, input)
		if _, err := io.Copy(os.Stdout, decoder); err != nil {
			return err
		}
		return nil
	},
}
