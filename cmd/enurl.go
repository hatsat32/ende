package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

var encodeall bool

// enurlCmd represents the enurl command
var enurlCmd = &cobra.Command{
	Use:   "enurl",
	Short: "URL encode given input",
	Run:   enurlCmdRun,
}

func enurlCmdRun(cmd *cobra.Command, args []string) {
	println(encodeall)
	fmt.Println(url.QueryEscape(strings.Join(args, " ")))
}

func init() {
	rootCmd.AddCommand(enurlCmd)
	enurlCmd.Flags().BoolVarP(&encodeall, "all", "a", false, "Encode all")
}
