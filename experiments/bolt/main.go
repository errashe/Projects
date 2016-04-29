package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("test"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		// err = b.Put([]byte("answer2"), []byte("qweqwe"))
		d := b.Get([]byte("answer"))
		fmt.Println(string(d))
		return nil
	})
}
