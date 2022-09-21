package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

// enurlCmd represents the enurl command
var enurlCmd = &cobra.Command{
	Use:   "enurl",
	Short: "URL encode given input",
	Run:   enurlCmdRun,
}

func enurlCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println(url.QueryEscape(strings.Join(args, " ")))
}

func init() {
	rootCmd.AddCommand(enurlCmd)
}
