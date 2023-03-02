package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/boltdb/bolt"
)

var (
	db   *bolt.DB
	once sync.Once
)

func GetDb() (*bolt.DB, error) {
	once.Do(func() {
		var err error
		db, err = bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			return nil
		})
	})
	return db, nil
}

func main() {

	// Example update read-write transaction
	db, err := GetDb()
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			return err
		}
		return b.Put([]byte("1"), []byte("Task 1"))
	})

	// Example reading transaction from the bucket
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		v := b.Get([]byte("1"))
		fmt.Println(string(v))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s", k, v)
			return nil
		})
		return nil
	})

	db.Close()
	//cmd.Execute()

}
