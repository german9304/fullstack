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
	age := 32

	user, _ := client.CreateUser(UserCreateInput{
		Email:    email,
		Password: password,
		Age:      int32(age),
	}).Exec(ctx)

	log.Printf("type => %T \n", user)
	log.Printf("Value => %v \n", user)
}
