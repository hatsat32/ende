package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

// deurlCmd represents the deurl command
var deurlCmd = &cobra.Command{
	Use:   "deurl",
	Short: "URL decode given input",
	Run:   deurlCmdRun,
}

func deurlCmdRun(cmd *cobra.Command, args []string) {
	out, _ := url.QueryUnescape(strings.Join(args, " "))
	fmt.Println(out)
}

func init() {
	rootCmd.AddCommand(deurlCmd)
}
