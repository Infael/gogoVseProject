package controller

import (
	"fmt"
	"net/http"
)

type Newsletter struct{}

func (n *Newsletter) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating newsletter...")
}

func (n *Newsletter) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing newsletters...")
}

func (n *Newsletter) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting newsletter by id...")
}

func (n *Newsletter) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating newsletter by id...")
}

func (n *Newsletter) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting newsletter by id...")
}

func (n *Newsletter) CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating new post for newsletter with id...")
}

func (n *Newsletter) GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting posts for newsletter with id...")
}
