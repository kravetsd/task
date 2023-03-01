package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do a task",
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		fmt.Printf("you have completted task \"%v\" \n", task)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
