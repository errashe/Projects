package main

import "fmt"
import "math"

// import "os"

var next int64 = 0

func setNext(nxt int64) {
	next = nxt
}

func reverse(num int64) int64 {
	return int64(uint32(math.MaxUint32 - num))
}

func rand() byte {
	next = next + 12512521*(1103515245^(reverse(next)%255))
	return (byte)(next/65536) % 255
}

func search(in []int, val int) []int {
	var res []int
	for i, n := range in {
		if n == val {
			res = append(res, i)
		}
	}
	return res
}

func cmp(in []int, cm []int) bool {
	for i := range in {
		if in[i] != cm[i] {
			return false
		}
	}
	return true
}

func main() {
	setNext(0)

	var for_find []int
	for i := 0; i < 10; i++ {
		for_find = append(for_find, int(rand()))
	}
	fmt.Println(for_find)

	var for_cmp []int
	cnt := 0
	for cnt < 100000 {
		num := int(rand())
		if num == for_find[0] {
			cnt++
		}
		for_cmp = append(for_cmp, num)
	}

	len_check := 5
	q := search(for_cmp, for_find[0])
	fmt.Println(len(for_cmp))
	for _, first := range q {
		if cmp(for_find[0:len_check], for_cmp[first:first+len_check]) == true {
			fmt.Printf("%d %d %d\n", for_find[0:len_check], for_cmp[first:first+len_check], first)
		}
	}

}
