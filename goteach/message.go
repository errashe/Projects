package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	Username string
	Message  string
	Created  string
}

var db *sql.DB

func Init() {
	db, _ = sql.Open("sqlite3", "./dev.db")
	db.Exec(`
		CREATE TABLE IF NOT EXISTS "messages" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"username" VARCHAR(64) NULL,
			"message" TEXT NULL,
			"created" DATE NULL
		);
	`, nil)
}

func AllMessages() ([]*Message, error) {
	rows, err := db.Query("SELECT username, message, created FROM messages ORDER BY created DESC LIMIT 15")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []*Message
	for rows.Next() {
		msg := new(Message)
		err := rows.Scan(&msg.Username, &msg.Message, &msg.Created)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return msgs, nil
}

func Insert(username, message, created string) {
	insertStmt, _ := db.Prepare("INSERT INTO 'messages' VALUES(NULL, ?, ?, ?)")
	insertStmt.Exec(username, message, created)
}
