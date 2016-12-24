package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/headzoo/surf"
)

func main() {
	a := surf.NewBrowser()

	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 1; i <= 100; i++ {
		db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("users"))
			if err != nil {
				return err
			}

			a.Open(fmt.Sprintf("http://vk.com/id%d", i))
			if a.Title() != "ВКонтакте" && a.Title() != "VK.com" {
				fmt.Println(a.Title())
				return b.Put([]byte(fmt.Sprintf("%d", i)), []byte(a.Title()))
			}

			return nil
		})
		time.Sleep(300 * time.Millisecond)
	}
}
