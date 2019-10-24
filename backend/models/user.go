package models

import (
	"time"
)

type User struct {
	id        string
	email     string
	name      string
	password  string
	age       int
	createdAt *time.Time
	posts     []*Post
	likes     []*Like
}

func NewUser(id, email, name, password string, age int, posts []*Post, likes []*Like) *User {
	createdAt := time.Now()
	return &User{id, email, name, password, age, &createdAt, posts, likes}
}

func (usr *User) ID() string {
	return usr.id
}

func (usr *User) Email() string {
	return usr.email
}

func (usr *User) Name() string {
	return usr.name
}

func (usr *User) Password() string {
	return usr.password
}

func (usr *User) Age() int {
	return usr.age
}

func (usr *User) CreatedAt() *time.Time {
	return usr.createdAt
}

func (usr *User) Posts() []*Post {
	return usr.posts
}

func (usr *User) Likes() []*Like {
	return usr.likes
}
