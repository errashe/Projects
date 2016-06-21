package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v1/?key=11B3469E9B4BB9C5349FA158D413FDF2&format=json&account_id=76561198052679375")
	if err != nil {
		println("HTTP.GET", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("IO.READALL", err.Error())
	}

	var m SteamApiMatches
	err = json.Unmarshal(body, &m)
	if err != nil {
		println("JSON.UNMARSHAL", err.Error())
	}

	fmt.Printf("%.f\n", m.Result.Matches[0].MatchID)
}
