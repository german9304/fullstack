package models

import (
	// "time"
	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

type Post struct {
	Id        string
	Text      string
	CreatedAt string
	UpdatedAt string
	Author    prisma.User
	Likes     []prisma.Like
}
