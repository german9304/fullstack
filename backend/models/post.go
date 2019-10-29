package models

import (
	// "time"
	prisma "github.com/german9304/fullstack-backend/prisma-client" 
)

type PostModel struct {
	post *prisma.Post
	author *prisma.User
	likes []prisma.Likes
}

// Init post model
func NewPostModel(post *prisma.Post, author *prisma.User, likes []prisma.Likes) *PostModel {
	return &PostModel{post, author, likes}
}


func (p *PostModel) Post() *prisma.Post {
	return p.post
}

func (p *PostModel) Author() *prisma.User {
	return p.author
}

func (p *PostModel) Likes() []prisma.Likes {
	return p.likes
}