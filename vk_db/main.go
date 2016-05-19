package main

import (
	"fmt"
	"strings"
)

func main() {
	res := VK(
		"users.get",
		fmt.Sprintf("user_ids=%s&fields=verified,sex,bdate,city,country", strings.Join(Range(0, 100, 1), ",")),
	)

	for _, user := range res.Metas {
		if user.Deactivated != "deleted" {
			fmt.Println(user.FirstName, user.LastName)
		}
	}

}
