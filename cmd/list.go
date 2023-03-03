package cmd

import (
	"encoding/binary"
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
	b := tx.Bucket([]byte(internal.BUCKETNAME))
	b.ForEach(func(k, v []byte) error {
		fmt.Printf("%v. %v\n", btoi(k), string(v))
		return nil
	})
	return nil
}

// btoi returns an int big endian representation of v.
func btoi(b []byte) uint64 {
	u := binary.BigEndian.Uint64(b)
	return u
}
