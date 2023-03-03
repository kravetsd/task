package cmd

import (
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/kravetsd/task/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%v\" into your list\n", task)
		internal.Db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
			if err != nil {
				return err
			}
			return b.Put([]byte(task), []byte(task))
		})
	},
}
