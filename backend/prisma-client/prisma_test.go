package prisma

import (
	"log"
	"testing"
	"context"
)

func TestUser(t *testing.T) {
	// "id","email","name","password","age","createdAt"
	client := New(nil)
	ctx := context.TODO()
	email := "John@mail.com"
	password := "293902122"

	user, _ := client.CreateUser(UserCreateInput{
		Email:    email,
		Password: password,
	}).Exec(ctx)

	log.Printf("type => %T \n", user)
	log.Printf("Value => %v \n", user.ID)

	// usrid := "2020202"

	newPost := client.CreatePost(PostCreateInput{
		Text:   "post one",
		Author: &UserCreateOneWithoutPostsInput{
			Connect: &UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
	})

	// if err != nil {
	// 	log.Printf("Error is %v \n", err)
	// }

	log.Println(newPost)

	postAuthor, _ := newPost.Author().Exec(ctx)
	log.Printf("Author: %v \n", postAuthor)
	log.Printf("Author type: %T \n", postAuthor)
}
