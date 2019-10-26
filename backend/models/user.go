package models

import (
	"time"
)

type User struct {
	Id        string
	Email     string
	Name      string
	Password  string
	Age       int
	CreatedAt *time.Time
	// Posts     []*Post
	// likes     []*Like
}


