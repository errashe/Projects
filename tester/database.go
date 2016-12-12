package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	Username string
	// Pwd      string
	Answers  int
	Ranswers int
}

type QuestionSet struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Count     int
	Text      string
	Variances map[string]string
}

type QuestionGet struct {
	Id        bson.ObjectId `json:"id" bson:"_id",omitempty"`
	RVariance []string
}

type QuestionInsert struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Text      string
	Variances map[string]string
	RVariance []string
}

var s *mgo.Session
var db string = "tester"

var users *mgo.Collection
var questions *mgo.Collection

func dbInit() {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Panic(err)
	}

	users = s.DB(db).C("users")
	questions = s.DB(db).C("questions")
}
