package main

import (
	"fmt"
	"html"
	"io"
	"net/url"

	"ende/internal/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(enhtmlCmd)
	rootCmd.AddCommand(dehtmlCmd)
	rootCmd.AddCommand(enurlCmd)
	rootCmd.AddCommand(deurlCmd)
}

// Helper to read all input as string, since HTML/URL functions operate on strings
func readInputString(args []string) (string, error) {
	input := cli.GetInput(args)
	data, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var enhtmlCmd = &cobra.Command{
	Use:   "enhtml [input]",
	Short: "Encode data to HTML Entities",
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := readInputString(args)
		if err != nil {
			return err
		}
		fmt.Print(html.EscapeString(str))
		return nil
	},
}

var dehtmlCmd = &cobra.Command{
	Use:   "dehtml [input]",
	Short: "Decode data from HTML Entities",
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := readInputString(args)
		if err != nil {
			return err
		}
		fmt.Print(html.UnescapeString(str))
		return nil
	},
}

var enurlCmd = &cobra.Command{
	Use:   "enurl [input]",
	Short: "URL Encode data",
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := readInputString(args)
		if err != nil {
			return err
		}
		fmt.Print(url.QueryEscape(str))
		return nil
	},
}

var deurlCmd = &cobra.Command{
	Use:   "deurl [input]",
	Short: "URL Decode data",
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := readInputString(args)
		if err != nil {
			return err
		}
		decoded, err := url.QueryUnescape(str)
		if err != nil {
			return err
		}
		fmt.Print(decoded)
		return nil
	},
}
