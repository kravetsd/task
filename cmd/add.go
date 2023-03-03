package cmd

import (
	"encoding/binary"
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
			b, err := tx.CreateBucketIfNotExists([]byte(internal.BUCKETNAME))
			if err != nil {
				return err
			}
			id, err := b.NextSequence()
			if err != nil {
				return err
			}
			return b.Put(itob(int(id)), []byte(task))
		})
	},
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
