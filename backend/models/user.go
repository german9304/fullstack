package models

import (
	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

type User struct {
	Id        string
	Email     string
	Name      string
	Password  string
	CreatedAt string
	UpdatedAt string
	Posts     []prisma.Post
	Likes     []prisma.Like
}
