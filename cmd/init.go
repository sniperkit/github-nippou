package cmd

import (
	"fmt"
	"os"

	"github.com/masutaka/github-nippou/lib"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize github-nippou settings",
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Init(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
