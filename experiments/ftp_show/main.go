package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
	"path/filepath"
	"strings"
)

type FI struct {
	Link string
	Name string
	Ext  string
	Dir  bool
}

func main() {
	var dir = flag.String("dir", ".", "Write directory name to share")
	var host = flag.String("host", "192.168.1.31:8080", "Write host and port to start server")
	flag.Parse()

	r := gin.Default()

	file, err := os.Open("auths.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	accs := gin.Accounts{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		creds := strings.Split(scanner.Text(), ":")
		accs[creds[0]] = creds[1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	authorized := r.Group(fmt.Sprintf("/%s", *dir), gin.BasicAuth(accs))

	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, *dir)
	})

	authorized.GET("/*dir", func(c *gin.Context) {
		path := fmt.Sprintf("%s%s", *dir, c.Param("dir"))

		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			c.Header("Content-Description", "File Transfer")
			c.Header("Content-Transfer-Encoding", "binary")
			c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", info.Name()))
			c.Header("Content-Type", "application/octet-stream")

			c.File(path)
			return
		}

		files, _ := filepath.Glob(fmt.Sprintf("%s/*", path))

		var fls []FI

		for _, file := range files {
			var isDir bool = false

			info, err := os.Stat(file)
			if err == nil && info.IsDir() {
				isDir = true
			}

			var ext string = "dir"
			if !isDir {
				parts := strings.Split(info.Name(), ".")
				if len(parts) > 1 {
					ext = parts[len(parts)-1]
				}
			}

			fls = append(fls, FI{strings.Replace(file, *dir, "", 1), info.Name(), ext, isDir})
		}

		c.HTML(200, "main.html", gin.H{"Root": fmt.Sprintf("/%s/", *dir), "root": *dir, "files": fls})
	})

	r.GET("/pic/:img", func(c *gin.Context) {
		c.File(fmt.Sprintf("./png/%s", c.Param("img")))
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(302, *dir)
	})

	r.Run(*host)

}
