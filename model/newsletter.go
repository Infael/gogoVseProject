package model

import "time"

type NewsletterAll struct {
	Id          uint64     `json:"id" validate:"required"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at" validate:"required"`
	Creator     uint64     `json:"creator" validate:"required"`
}

type NewsletterCreate struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Creator     uint64 `json:"creator"`
}

type NewsletterUpdate struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type NewsletterAllList struct {
	Newsletters []NewsletterAll `json:"newsletters"`
}
