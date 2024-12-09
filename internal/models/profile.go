package models

import "time"

type Profile struct {
	Id           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	HashPassword string    `json:"hashpassword"`
	Status       bool      `json:"status"`
	AccessLevel  int       `json:"accesslevel"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	CreatedAt    time.Time `json:"createdat"`
}

type ProfileCheck struct {
	Nickname     string `json:"nickname"`
	HashPassword string `json:"hashpassword"`
}
