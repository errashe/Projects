package main

import (
	"flag"
	"fmt"
	"github.com/headzoo/surf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type User struct {
	gorm.Model
	Uid  uint
	Name string
}

func main() {
	var start = flag.Int("start", 1, "Set starting number")
	var end = flag.Int("end", 100, "Set ending number")

	flag.Parse()

	a := surf.NewBrowser()

	db, err := gorm.Open("postgres", "host=localhost user=e4stw00d dbname=e4stw00d sslmode=disable password=")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	for i := *start; i <= *end; i++ {
		a.Open(fmt.Sprintf("http://vk.com/id%d", i))
		if a.Title() != "ВКонтакте" && a.Title() != "VK.com" {
			fmt.Printf("%d - %s\n", i, a.Title())
			db.Create(&User{Uid: (uint)(i), Name: a.Title()})
		}
		time.Sleep(300 * time.Millisecond)
	}

	var users []User
	db.Find(&users)

	fmt.Println(len(users))
}
