package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"html/template"
)

var (
	funcs []template.FuncMap
)

func isUser(c *gin.Context) bool {
	return s(c).Get("user") != nil
}

func User(c *gin.Context) string {
	if isUser(c) {
		return s(c).Get("user").(string)
	} else {
		return "none"
	}
}

func init() {
	funcs = append(funcs, template.FuncMap{
		"isUser": isUser,
		"User":   User,
	})
}
