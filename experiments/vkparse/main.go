package main

import (
	"flag"
	. "fmt"
	"github.com/Jeffail/gabs"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type User struct {
	Id        float64 `csv:"id"`
	FirstName string  `csv:"firstname"`
	LastName  string  `csv:"lastname"`
	City      string  `csv:"city"`
	Bdate     string  `csv:"bdate"`
}

type Users []User

func getHundred(start int) Users {
	var re []string
	for i := start; i < start+100; i++ {
		re = append(re, Sprintf("%d", i))
	}

	ids := strings.Join(re, ",")
	fields := "verified,sex,bdate,city,has_mobile,contacts,site,connections"

	url := Sprintf("https://api.vk.com/method/users.get?user_ids=%s&fields=%s&v=5.62", ids, fields)

	res, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	vkData, err := gabs.ParseJSON(bytes)
	if err != nil {
		Println(err)
		os.Exit(0)
	}

	users, err := vkData.S("response").Children()
	if err != nil {
		Println(err)
		os.Exit(0)
	}

	var retUsers Users

	for _, user := range users {
		if user.Path("deactivated").Data() == nil {
			retUser := User{}

			retUser.Id = user.Path("id").Data().(float64)
			if firstnameraw := user.Path("first_name").Data(); firstnameraw != nil {
				retUser.FirstName = firstnameraw.(string)
			}
			if lastnameraw := user.Path("last_name").Data(); lastnameraw != nil {
				retUser.LastName = lastnameraw.(string)
			}
			if cityraw := user.Path("city.title").Data(); cityraw != nil {
				retUser.City = cityraw.(string)
			}
			if bdateraw := user.Path("bdate").Data(); bdateraw != nil {
				retUser.Bdate = bdateraw.(string)
			}

			retUsers = append(retUsers, retUser)
		}
	}

	return retUsers
}

func main() {
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	// last := 80000000
	start := 0
	end := 80000000

	for i := start; i < end; i += 100 {
		users := getHundred(i)
		if i == start {
			gocsv.MarshalFile(users, clientsFile)
		} else {
			gocsv.MarshalWithoutHeaders(users, clientsFile)
		}
		Println(len(users), users[0].Id)
		time.Sleep(1 * time.Second)
	}
}
