package models

import "time"

type Post struct {
	id        string
	text      string
	createdAt time.Time
	updatedAt time.Time
	author    string
	likes     []*Like
}
