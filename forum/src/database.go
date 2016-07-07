package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Section struct {
	ID   int    `db:"ID"`
	NAME string `db:"NAME"`
}

type Theme struct {
	ID         int    `db:"ID"`
	SECTION_ID int    `db:"SECTION_ID"`
	NAME       string `db:"NAME"`
}

type Message struct {
	ID       int    `db:"ID"`
	THEME_ID int    `db:"THEME_ID"`
	USER_ID  int    `db:"USER_ID"`
	MESSAGE  string `db:"MESSAGE"`
}

func dbInit() *sqlx.DB {
	schema := `
	CREATE TABLE IF NOT EXISTS 'users' (
		'ID' integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		'LOGIN' varchar(255) NOT NULL,
		'PASSWORD' varchar(255) NOT NULL,
		'USERNAME' varchar(255) NOT NULL,
		'EMAIL' varchar(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS 'sections' (
		'ID' integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		'NAME' varchar(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS 'themes' (
		'ID' integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		'SECTION_ID' integer NOT NULL,
		'NAME' varchar(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS 'messages' (
		'ID' integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		'THEME_ID' integer NOT NULL,
		'USER_ID' integer NOT NULL,
		'MESSAGE' text NOT NULL
	);
	`

	db, err := sqlx.Connect("sqlite3", "test.db")
	ehandle(err)
	db.MustExec(schema)

	return db
}
