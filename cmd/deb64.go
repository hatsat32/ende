package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// deb64Cmd represents the deb64 command
var deb64Cmd = &cobra.Command{
	Use:   "deb64",
	Short: "Decode base64 given input",
	Run:   deb64Run,
}

func deb64Run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("[!] Give one and only one argument!")
		return
	}
	decoded, _ := base64.StdEncoding.DecodeString(args[0])
	fmt.Println(string(decoded))
}

func init() {
	rootCmd.AddCommand(deb64Cmd)
}
