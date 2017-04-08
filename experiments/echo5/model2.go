package main

import . "fmt"
import "encoding/json"

func work2(rp ReqParams) {
	defer wg.Done()

	var m Matrix
	json.Unmarshal([]byte(rp.Par2), &m)

	Println(len(m), len(m[0]))
}
