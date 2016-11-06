package main

import (
	"crypto/md5"
	"fmt"
	"github.com/asdine/storm"
	// "github.com/asdine/storm/q"
	"github.com/gorilla/websocket"
)

type User struct {
	ID           int    `storm:"id,increment"`
	Username     string `storm:"unique"`
	PasswordHash string
}

type Message struct {
	ID       int `storm:"id,increment"`
	Username string
	Body     string
}

var db *storm.DB

func initDB() {
	db, _ = storm.Open("my.db")
}

func addUser(username, password string) {
	h := md5.New()
	h.Write([]byte(password))
	user := User{
		Username:     username,
		PasswordHash: fmt.Sprintf("%x", h.Sum(nil)),
	}

	error_handler(db.Save(&user))
}

func checkUser(username, password string) bool {
	h := md5.New()
	h.Write([]byte(password))
	user := User{}
	err := db.One("Username", username, &user)
	if err != nil {
		return false
	}

	if user.PasswordHash == fmt.Sprintf("%x", h.Sum(nil)) {
		return true
	}
	return false
}

func getUser(id int) User {
	user := User{}
	error_handler(db.One("ID", id, &user))
	return user
}

func addMessage(username, body string) {
	message := Message{
		Username: username,
		Body:     body,
	}

	broadcastAll([]byte(fmt.Sprintf("%s: %s", username, body)))

	err := db.Save(&message)
	error_handler(err)
}

func getLastFifteenMessages(conn *websocket.Conn) {
	var messages []Message
	error_handler(db.AllByIndex("ID", &messages, storm.Limit(15), storm.Reverse()))
	// error_handler(db.Select(q.True()).Reverse().Limit(14).OrderBy("ID").Find(&messages))

	// for _, msg := range messages {
	// 	broadcastOne(conn, []byte(fmt.Sprintf("%s: %s", msg.Username, msg.Body)))
	// }

	for i := len(messages) - 1; i >= 0; i-- {
		msg := messages[i]
		broadcastOne(conn, []byte(fmt.Sprintf("%s: %s", msg.Username, msg.Body)))
	}
}
