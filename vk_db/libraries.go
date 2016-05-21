package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var token string = "18125f3dea048c0dfc3b75b906691ee20a603e56ba0cdaaefcdca43733602d1e96b07c7a199e3077a175f"

func VK(method, params string) Resp {
	time.Sleep(200 * time.Millisecond)
	link := fmt.Sprintf("https://api.vk.com/method/%s?%s&access_token=%s", method, params, token)
	// fmt.Println(link)
	resp, err := http.Get(link)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)

	var ret Resp

	err = json.Unmarshal(str, &ret)
	if err != nil {
		log.Println(err.Error())
	}

	return ret
}

func Range(start, end, step int) []string {
	var res []string
	for i := start; i < end; i += step {
		res = append(res, fmt.Sprintf("%d", i))
	}
	return res
}

func OpenDB(name string) *bolt.DB {
	db, err := bolt.Open(name, 0600, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func WriteData(db *bolt.DB, users Resp) {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		for _, user := range users.Metas {
			if user.Deactivated == "deleted" || user.Deactivated == "banned" {
				continue
			}
			d, err := json.Marshal(user)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = b.Put([]byte(fmt.Sprintf("%d", time.Now().UnixNano())), d)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		return nil
	})
}
