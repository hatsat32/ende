package main

import (
	"encoding/base64"
	"io"
	"os"

	"ende/internal/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(enb64Cmd)
	rootCmd.AddCommand(deb64Cmd)
}

var enb64Cmd = &cobra.Command{
	Use:   "enb64 [input]",
	Short: "Encode data to Base64",
	RunE: func(cmd *cobra.Command, args []string) error {
		input := cli.GetInput(args)
		encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
		if _, err := io.Copy(encoder, input); err != nil {
			return err
		}
		return encoder.Close()
	},
}

var deb64Cmd = &cobra.Command{
	Use:   "deb64 [input]",
	Short: "Decode data from Base64",
	RunE: func(cmd *cobra.Command, args []string) error {
		input := cli.GetInput(args)
		decoder := base64.NewDecoder(base64.StdEncoding, input)
		if _, err := io.Copy(os.Stdout, decoder); err != nil {
			return err
		}
		return nil
	},
}
