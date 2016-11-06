package main

import (
	"github.com/gorilla/websocket"

	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func error_handler(err error) {
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
}

func control() {
	for {
		var str string
		fmt.Scanln(&str)
		s := strings.Split(str, ",")

		switch s[0] {
		case "au":
			addUser(s[1], s[2])
		case "am":
			addMessage(s[1], s[2])
		case "gm":
			id, err := strconv.Atoi(s[1])
			error_handler(err)
			fmt.Printf("%v\n", getUser(id))
		case "cu":
			fmt.Println(checkUser(s[1], s[2]))
		}
	}
}

func broadcastOne(conn *websocket.Conn, msg []byte) {
	conn.WriteMessage(1, msg)
}

func broadcastAll(msg []byte) {
	for _, c := range conns {
		broadcastOne(c, msg)
	}
}

func loadTemplates() *template.Template {
	var templates = template.New("")

	funcMap := template.FuncMap{
		"isSessionIsset": func(r *http.Request) bool {
			session, _ := store.Get(r, "main")
			if session.Values["nick"] == nil {
				return false
			}
			return true
		},
	}

	templates.Funcs(funcMap)

	for _, path := range AssetNames() {
		bytes, err := Asset(path)
		error_handler(err)
		templates.New(path).Parse(string(bytes))
	}

	return templates
}
