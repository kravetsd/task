package internal

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type Task struct {
	Id    int
	Value string
}

func NewTask(id int) *Task {
	task := &Task{
		Id: id,
	}
	return task
}

func (t *Task) DoTask(tx *bolt.Tx) error {
	bucket := tx.Bucket([]byte(BUCKETNAME))
	err := bucket.Delete(itob(t.Id))
	// TODO: need to check if task does not exists
	if err != nil {
		log.Println("Error while deleting the task: ", err)
		return err
	}
	fmt.Printf("you have completted task \"%v\" \n", t.Id)
	return err
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
