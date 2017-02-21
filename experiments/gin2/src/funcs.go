package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"html/template"
)

var (
	funcs template.FuncMap
)

func test(c *gin.Context) string {
	sess := se(c).Get("test")
	if sess != nil {
		return sess.(string)
	}
	return "none"
}

func init() {
	funcs = template.FuncMap{
		"test": test,
	}
}
