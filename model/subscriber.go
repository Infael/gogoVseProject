package model

type Subscribe struct {
	Email string `json:"email" validate:"required"`
}

type SubscriberAll struct {
	Id    uint64 `json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type SubscriberAllList struct {
	Subscribers []SubscriberAll `json:"subscribers" validate:"required"`
}
