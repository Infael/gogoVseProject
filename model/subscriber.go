package model

type Subscriber struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
}

type SubscriberList struct {
	Subscribers []Subscriber `json:"users"`
}
