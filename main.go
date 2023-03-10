package main

import (
	"log"

	"github.com/kravetsd/task/cmd"
	"github.com/kravetsd/task/internal"
)

func main() {

	// Example update read-write transaction
	DataBase, err := internal.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer DataBase.Close()

	// Example reading transaction from the bucket
	// DataBase.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("Tasks"))
	// 	v := b.Get([]byte("1"))
	// 	fmt.Println(string(v))
	// 	return nil
	// })

	// DataBase.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("Tasks"))
	// 	b.ForEach(func(k, v []byte) error {
	// 		fmt.Printf("key=%s, value=%s", k, v)
	// 		return nil
	// 	})
	// 	return nil
	// })
	cmd.Execute()

}
