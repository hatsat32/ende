package cmd

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// enb64Cmd represents the enb64 command
var enb64Cmd = &cobra.Command{
	Use:   "enb64",
	Short: "Encode base64 given input",
	Run:   enb64Run,
}

func enb64Run(cmd *cobra.Command, args []string) {
	encoded := base64.StdEncoding.EncodeToString([]byte(strings.Join(args, " ")))
	fmt.Println(encoded)
}

func init() {
	rootCmd.AddCommand(enb64Cmd)
}
