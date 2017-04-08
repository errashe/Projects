package main

import "time"

type Settings struct {
	Cores   int `query:"cores"`
	Threads int `query:"threads"`
}

type Cell int
type Row []Cell
type Matrix []Row
type MatrixCalc struct {
	M1 Matrix
	M2 Matrix
}

type ReqParams struct {
	Work int    `query:"work"`
	Par1 int    `query:"par1"`
	Par2 string `form:"data"`
}

type Result struct {
	Res  float64       `json:"Res"`
	Time time.Duration `json:"Time"`
}

type Results []Result
