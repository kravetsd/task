package internal

import (
	"fmt"
	"log"
	"sync"

	"github.com/boltdb/bolt"
)

var (
	Db   *bolt.DB
	once sync.Once
)

const (
	BUCKETNAME string = "Tasks"
)

// type Task struct {
// 	Key   int
// 	Value string
// }

// type TaskList []Task

func GetDb() (*bolt.DB, error) {
	once.Do(func() {
		var err error
		Db, err = bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		Db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(BUCKETNAME))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			return nil
		})
	})
	return Db, nil
}
