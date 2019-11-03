package models

import (
	"time"


	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

type Like struct {
	Id        string
	Quantity  int
	User      prisma.User
	Post      prisma.Post
	CreatedAt time.Time
	UpdatedAt time.Time
	
}
