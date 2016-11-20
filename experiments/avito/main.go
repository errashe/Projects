package main

import (
	"fmt"
	"github.com/headzoo/surf"
	"net/http"
	"os"
	"strings"
)

var client http.Client
var req *http.Request
var res *http.Response

func main() {
	a := surf.NewBrowser()

	a.Open("https://www.avito.ru/registration")

	var cookies []string
	for _, cookie := range a.SiteCookies() {
		cookies = append(cookies, fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}

	req, _ = http.NewRequest("GET", a.Images()[0].URL.String(), nil)
	req.Header.Add("Referer", "https://www.avito.ru/registration")
	req.Header.Add("Cookie", strings.Join(cookies, "; "))

	res, _ = client.Do(req)

	fout, _ := os.Create("test.jpg")
	defer fout.Close()

	buffer := make([]byte, 5*1024*1024)

	n, _ := res.Body.Read(buffer)
	fout.Write(buffer[:n])

	var captcha string
	fmt.Scanln(&captcha)

	form, _ := a.Form("form.form")
	form.Input("email", "kirilloff@mail.wtf")
	form.Input("name", "Кирилов Владимир")
	form.Input("phone", "+7-919-596-69-79")
	form.Input("type", "0")
	form.Input("enablePro", "0")
	form.Input("password", "cvyqjkeg45")
	form.Input("confirm", "cvyqjkeg45")
	form.Input("captcha", captcha)
	form.Submit()

	iout, _ := os.Create("index.html")
	defer iout.Close()
	a.Download(iout)
}
