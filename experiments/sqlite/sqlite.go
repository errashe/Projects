package main

import "fmt"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

// TestItem qwe
type TestItem struct {
	ID    string
	Name  string
	Phone string
}

// InitDB qwe
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

// CreateTable qwe
func CreateTable(db *sql.DB) {
	// create table if not exists
	sqlTable := `
	CREATE TABLE IF NOT EXISTS items(
		ID TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		Phone TEXT,
		InsertedDatetime DATETIME
		);
		`

	_, err := db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
}

// StoreItem qwe
func StoreItem(db *sql.DB, items []TestItem) {
	sqlAdditem := `
		INSERT OR REPLACE INTO items(
			ID,
			Name,
			Phone,
			InsertedDatetime
			) values(?, ?, ?, CURRENT_TIMESTAMP)
			`

	stmt, err := db.Prepare(sqlAdditem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.ID, item.Name, item.Phone)
		if err2 != nil {
			panic(err2)
		}
	}
}

// ReadItem qwe
func ReadItem(db *sql.DB) []TestItem {
	sqlReadall := `
			SELECT ID, Name, Phone FROM items
			ORDER BY datetime(InsertedDatetime) DESC
			`

	rows, err := db.Query(sqlReadall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []TestItem
	for rows.Next() {
		item := TestItem{}
		err2 := rows.Scan(&item.ID, &item.Name, &item.Phone)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func main() {

	const dbpath = "foo.db"

	db := InitDB(dbpath)
	defer db.Close()
	CreateTable(db)

	items := []TestItem{
		TestItem{"1", "A", "213"},
		TestItem{"2", "B", "214"},
	}
	StoreItem(db, items)

	readItems := ReadItem(db)
	fmt.Println(readItems)

	items2 := []TestItem{
		TestItem{"153", "HAS", "3141"},
	}
	StoreItem(db, items2)

	readItems2 := ReadItem(db)
	fmt.Println(readItems2)
}
