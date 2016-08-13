package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
)

func getPosts(c *gin.Context) {
	posts, _ := ioutil.ReadDir("./posts")
	var json_posts []string
	for _, post := range posts {
		name := strings.Replace(post.Name(), ".txt", "", -1)
		json_posts = append(json_posts, name)
	}
	c.JSON(200, gin.H{
		"posts": json_posts,
	})
}

func getPost(c *gin.Context) {
	file, _ := ioutil.ReadFile(fmt.Sprintf("posts/%s.txt", c.Param("name")))
	post := string(file)
	c.JSON(200, gin.H{
		"title": c.Param("name"),
		"body":  post,
	})
}

func savePost(c *gin.Context) {
	ioutil.WriteFile(fmt.Sprintf("posts/%s.txt", c.PostForm("title")), []byte(c.PostForm("body")), 0644)
	c.JSON(200, gin.H{
		"title": c.PostForm("title"),
		"body":  c.PostForm("body"),
	})
}

func updatePost(c *gin.Context) {
	ioutil.WriteFile(fmt.Sprintf("posts/%s.txt", c.Param("name")), []byte(c.PostForm("body")), 0644)
	c.JSON(200, gin.H{
		"body": c.PostForm("body"),
	})
}

func deletePost(c *gin.Context) {
	os.Remove(fmt.Sprintf("posts/%s.txt", c.Param("name")))
	c.JSON(200, gin.H{
		"title": c.Param("name"),
	})
}

func main() {
	r := gin.Default()

	r.GET("/", getPosts)
	r.GET("/:name", getPost)
	r.POST("/", savePost)
	r.PUT("/:name", updatePost)
	r.DELETE("/:name", deletePost)

	r.Run(":3001")
}
