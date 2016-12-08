package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"reflect"
	"sort"
)

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	session := sessions.Default(c)
	user := &User{}
	session.Delete("userID")

	err := users.Find(bson.M{
		"username": username,
		"pwd":      fmt.Sprintf("%x", md5.Sum([]byte(password))),
	}).One(&user)

	var message string
	if err == nil {
		session.Set("userID", user.Username)
		message = "DONE"
	} else {
		message = "ERROR"
	}
	session.Save()
	c.JSON(200, gin.H{"message": message, "username": user.Username})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("userID")
	session.Save()

	c.JSON(200, gin.H{"message": "DONE"})
}

func currentUser(c *gin.Context) {
	session := sessions.Default(c)

	user := &User{}
	err := users.Find(bson.M{"username": session.Get("userID")}).One(&user)

	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(200, gin.H{"message": "ERROR"})
	}
}

func question(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("userID")

	q := QuestionSet{}
	n, _ := questions.Count()
	questions.Find(bson.M{}).Limit(-1).Skip(rand.Intn(n)).One(&q)
	users.Update(bson.M{"username": uid}, bson.M{"$inc": bson.M{"answers": 1}})

	c.JSON(200, q)
}

func answer(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("userID")

	q := QuestionGet{}
	questions.FindId(bson.ObjectIdHex(c.PostForm("id"))).One(&q)

	var variance []string
	json.Unmarshal([]byte(c.PostForm("variance")), &variance)
	sort.Strings(variance)

	var message string
	if reflect.DeepEqual(variance, q.RVariance) {
		if uid != nil {
			users.Update(bson.M{"username": uid}, bson.M{"$inc": bson.M{"ranswers": 1}})
		}
		message = "RIGHT"
	} else {
		message = "WRONG"
	}
	c.JSON(200, gin.H{"message": message})
}

func test(c *gin.Context) {
	q := QuestionInsert{}
	q.Text = "OAMSG)IASNP AS A"
	q.Variances = make(map[string]string)
	q.Variances["v1"] = "v11"
	q.Variances["v2"] = "v22"
	q.Variances["v3"] = "v33"
	q.Variances["v4"] = "v44"
	q.RVariance = []string{"v2"}

	questions.Insert(q)
	c.String(200, "DONE")
}

func main() {
	dbInit()
	defer s.Close()

	store := sessions.NewCookieStore([]byte("secret_word"))

	r := gin.Default()
	r.Use(sessions.Sessions("mysess", store))
	r.LoadHTMLFiles("index.html")

	r.GET("/", index)
	r.POST("/login", login)
	r.POST("/logout", logout)
	r.GET("/user", currentUser)
	r.GET("/question", question)
	r.POST("/answer", answer)
	r.GET("/test", test)

	r.Run(":3000")
}
