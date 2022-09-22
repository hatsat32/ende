/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// dejwtCmd represents the dejwt command
var dejwtCmd = &cobra.Command{
	Use:   "dejwt",
	Short: "Decode JWT Token",
	Run:   dejwtCmdRun,
}

func dejwtCmdRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("[!] Give one and only one argument!")
		return
	}

	token := strings.Split(args[0], ".")
	if len(token) != 3 {
		fmt.Println("[!] No valid JWT")
		return
	}

	tokenHeader, err := base64.RawURLEncoding.DecodeString(token[0])
	if err != nil {
		fmt.Println("[!] No valid JWT")
		return
	}

	tokenBody, err := base64.RawURLEncoding.DecodeString(token[1])
	if err != nil {
		fmt.Println("[!] No valid JWT")
		return
	}

	fmt.Println("[+] Header: ", string(tokenHeader))
	fmt.Println("[+] Body: ", string(tokenBody))
}

func init() {
	rootCmd.AddCommand(dejwtCmd)
}
