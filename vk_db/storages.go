package main

type Resp struct {
	Metas []Meta `json:"response"`
}

type Meta struct {
	UID         int    `json:"uid"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Sex         int    `json:"sex"`
	City        int    `json:"city,omitempty"`
	Country     int    `json:"country,omitempty"`
	Verified    int    `json:"verified,omitempty"`
	Deactivated string `json:"deactivated,omitempty"`
	Bdate       string `json:"bdate,omitempty"`
}
