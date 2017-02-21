package main

import (
	"gopkg.in/mgo.v2"
)

var (
	err     error
	session *mgo.Session
)

func init() {
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

}
