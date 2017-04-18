package dispatcher

import (
	. "../types"
	"time"
)

type Request1 struct {
	M1 Matrix
	M2 Matrix
}

type Response1 struct {
	Answer Matrix
	Time   time.Duration
}
