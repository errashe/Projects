package main

type Friend struct {
	UID          int    `json:"uid"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Online       int    `json:"online"`
	UserID       int    `json:"user_id"`
	Bdate        string `json:"bdate,omitempty"`
	Twitter      string `json:"twitter,omitempty"`
	Instagram    string `json:"instagram,omitempty"`
	Skype        string `json:"skype,omitempty"`
	Lists        []int  `json:"lists,omitempty"`
	Facebook     string `json:"facebook,omitempty"`
	FacebookName string `json:"facebook_name,omitempty"`
	Deactivated  string `json:"deactivated,omitempty"`
}

type Profile struct {
	UID       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Counters  struct {
		Albums        int `json:"albums"`
		Videos        int `json:"videos"`
		Audios        int `json:"audios"`
		Notes         int `json:"notes"`
		Photos        int `json:"photos"`
		Friends       int `json:"friends"`
		OnlineFriends int `json:"online_friends"`
		MutualFriends int `json:"mutual_friends"`
		Followers     int `json:"followers"`
		Subscriptions int `json:"subscriptions"`
		Pages         int `json:"pages"`
	} `json:"counters"`
}

type Friends struct {
	Response []Friend `json:"response"`
}

type Profiles struct {
	Response []Profile `json:"response"`
}
