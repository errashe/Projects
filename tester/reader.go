package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dbInit()
	// defer s.Close()

	f, _ := os.Open("test.csv")
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'

	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		res_s := make(map[string]string)
		for i, s := range record[3:11] {
			if len(s) > 0 {
				res_s[fmt.Sprintf("v%d", i+1)] = s
			}
		}

		q := QuestionInsert{}
		q.Text = record[2]
		q.Variances = res_s
		q.RVariance = strings.Split(record[11], ",")

		if len(record[11]) != 0 {
			questions.Insert(q)
		}
	}

	fmt.Println("DONE")
}
