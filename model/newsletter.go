package model

import "time"

type Newsletter struct {
	Id          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Posts       []Post     `json:"posts"`
	Subscribers []User     `json:"subscribers"`
	CreatedAt   *time.Time `json:"created_at"`
	Creator     uint64     `json:"creator"`
}
