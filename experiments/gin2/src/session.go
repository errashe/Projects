package main

import (
	. "fmt"
	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	store sessions.CookieStore
)

func se(c *gin.Context) sessions.Session { return sessions.Default(c) }

func init() {
	Println("session loaded")

	store = sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("gin", store))
}
