package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	db := OpenDB("users.db")
	defer db.Close()

	for i := 0; i < 100000/100; i++ {
		res := VK(
			"users.get",
			fmt.Sprintf("user_ids=%s&fields=verified,sex,bdate,city,country", strings.Join(Range(i*100, i*100+100, 1), ",")),
		)

		WriteData(db, res)
		fmt.Println(i * 100)
	}

	fmt.Println("DONE")

	fmt.Println(time.Since(t))

}
