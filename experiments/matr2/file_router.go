package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func file_router() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/", http.Dir("./assets"))

	return r
}
