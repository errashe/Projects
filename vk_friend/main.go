package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var token string = "18125f3dea048c0dfc3b75b906691ee20a603e56ba0cdaaefcdca43733602d1e96b07c7a199e3077a175f"

func VK(method, params string) []byte {
	link := fmt.Sprintf("https://api.vk.com/method/%s?%s&access_token=%s", method, params, token)
	// fmt.Println(link)
	time.Sleep(300 * time.Millisecond)
	resp, err := http.Get(link)
	if err != nil {
		log.Panic(err.Error())
	}
	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)

	return str
}

func RemoveDuplicates(xs *[]int) {
	found := make(map[int]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func getProfiles(uid int) Profile {
	var ret Profiles
	f := VK("users.get", fmt.Sprintf("user_id=%d&fields=counters", uid))
	json.Unmarshal(f, &ret)
	return ret.Response[0]
}

func getFriends(uid int) Friends {
	var ret Friends
	f := VK("friends.get", fmt.Sprintf("user_id=%d&fields=bdate,connections", uid))
	json.Unmarshal(f, &ret)
	return ret
}

func main() {
	db, err := bolt.Open("friends.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if len(os.Args) < 2 {
		var start int = 5387082
		var all_friends []Friend

		f := getFriends(start)
		for _, friend := range f.Response {
			temp := getFriends(friend.UID)
			all_friends = append(all_friends, temp.Response...)
		}

		fmt.Println(len(all_friends))

		for _, f := range all_friends {
			db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte("friends"))
				if err != nil {
					return err
				}
				encoded, err := json.Marshal(f)
				if err != nil {
					return err
				}
				return b.Put([]byte(fmt.Sprintf("%d", time.Now().UnixNano())), encoded)
			})
		}
	} else {
		var res []int
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("friends"))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				var f Friend
				json.Unmarshal(v, &f)
				if strings.Contains(f.Bdate, "1990") && strings.Contains(f.FirstName, "Алек") {
					res = append(res, f.UID)
				}
			}

			return nil
		})

		RemoveDuplicates(&res)
		for _, id := range res {
			u := getProfiles(id)
			if u.Counters.Friends < 50 {
				fmt.Println(u.UID)
			}
		}
	}

}
