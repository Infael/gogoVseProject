package model

import "time"

type Post struct {
	Id         uint64     `json:"id"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	CreatedAt  *time.Time `json:"created_at"`
	Newsletter uint64     `json:"newsletter_id"`
}
