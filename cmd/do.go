package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do a task")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
