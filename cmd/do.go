package cmd

import (
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/kravetsd/task/internal"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do a task",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		t := internal.NewTask(int(id))
		internal.Db.Update(func(tx *bolt.Tx) error {
			return t.DoTask(tx)
		})

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
