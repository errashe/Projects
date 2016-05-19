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
	link := fmt.Sprintf("https://api.vk.com/method/%s?%s&access_token=%s", method, params, token)
	// fmt.Println(link)
	resp, err := http.Get(link)
	if err != nil {
		log.Panic(err.Error())
	}
	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)

	var ret Resp

	err = json.Unmarshal(str, &ret)
	if err != nil {
		log.Panic(err.Error())
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

func WriteToBase() {
	db, err := bolt.Open("users.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		return b.Put([]byte(fmt.Sprintf("%f", time.Now().UnixNano())), []byte("Hello"))
	})
}
