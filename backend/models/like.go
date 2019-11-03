package models

import (
	"time"


	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

type Like struct {
	Id        string
	Quantity  int
	user      prisma.User
	post      prisma.Post
	createdAt time.Time
	updatedAt time.Time
	
}
