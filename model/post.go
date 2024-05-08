package model

import "time"

type PostAll struct {
	Id           uint64     `json:"id" validate:"required"`
	Title        string     `json:"title" validate:"required"`
	Body         string     `json:"body" validate:"required"`
	CreatedAt    *time.Time `json:"created_at" validate:"required"`
	NewsletterId uint64     `json:"newsletter_id" validate:"required"`
}

type PostUpdate struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type PostAllList struct {
	Posts []PostAll `json:"posts" validate:"required"`
}
