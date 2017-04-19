package dispatcher

import (
	. "../types"
	"time"
)

type Request0 struct {
	M1 Pi
}

type Response0 struct {
	Answer float64
	Time   time.Duration
}

type Request1 struct {
	M1 Matrix
	M2 Matrix
}

type Response1 struct {
	Answer Matrix
	Time   time.Duration
}

type Request2 struct {
	M1 Matrix
	M2 VerticalVector
}

type Response2 struct {
	Answer VerticalVector
	Time   time.Duration
}

type Request3 struct {
	M1 VerticalVector
	M2 HorizontalVector
}

type Response3 struct {
	Answer Matrix
	Time   time.Duration
}

type Request4 struct {
	M1 Brute
}

type Response4 struct {
	Answer []UserData
	Time   time.Duration
}
