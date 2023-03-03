package internal

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

type Task struct {
	Id    int
	Value string
}

func (t *Task) doTask(tx *bolt.Tx) error {
	bucket := tx.Bucket([]byte(BUCKETNAME))
	return bucket.Delete(itob(t.Id))
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi returns an int representation of byte.
func btoi(b []byte) uint64 {
	u := binary.BigEndian.Uint64(b)
	return u
}

// TODO: finilize this method to create task instance from bolt db
// func GetTask(id string) (*Task, error) {
// 	task := &Task{}
// 	i, err := strconv.ParseInt(id, 10, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	task.Id = int(btoi([]byte(id)))

// }
