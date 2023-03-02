package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/kravetsd/task/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Run:   runListTasks,
}

func runListTasks(cmd *cobra.Command, args []string) {
	internal.Db.View(listTasks)
}

func listTasks(tx *bolt.Tx) error {
	var tasks []string
	b := tx.Bucket([]byte(internal.BUCKETNAME))
	b.ForEach(func(k, v []byte) error {
		tasks = append(tasks, string(v))
		return nil
	})
	for i, t := range tasks {
		fmt.Printf("%d. %v", i+1, t)
	}
	return nil
}
