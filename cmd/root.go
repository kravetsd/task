package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI simple task manager",
	Long:  "task CLI tool allows you to perform CRUD operations on your day by day life TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {

	// To make a "help" flag to be hidden from the help menu we need to define it first
	rootCmd.PersistentFlags().BoolP("help", "h", false, "help for task")
	rootCmd.PersistentFlags().Lookup("help").Hidden = true
	// disabling "help" command from the help menu
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	// Disabling "completion" command from the help menu
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
