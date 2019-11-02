package fullstack_backend

import (
	"log"
	"testing"

	prisma "github.com/german9304/fullstack-backend/prisma-client"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

type FullstackSuiteQuery struct {
	suite.Suite
	queryUsers []prisma.UserCreateInput
	clientGraphql *graphql.Client
}

func (fs *FullstackSuiteQuery) SetupSuite() {

	fs.clientGraphql = graphql.NewClient("http://localhost:8000/")

	fs.queryUsers = []prisma.UserCreateInput{
		prisma.UserCreateInput{
			Email:    "pepe@mail.com",
			Name:     "Pepe",
			Password: "pepe1234",
		},
		prisma.UserCreateInput{
			Email:    "jimmy@mail.com",
			Name:     "Jimmy",
			Password: "jim1234",
		},
		prisma.UserCreateInput{
			Email:    "miguel@mail.com",
			Name:     "Miguel",
			Password: "miguel1234",
		},
		prisma.UserCreateInput{
			Email:    "kendrick@mail.com",
			Name:     "Kendrick",
			Password: "kendrick1234",
		},
	}

	for i := 0; i < len(fs.queryUsers); i++ {
		user := fs.queryUsers[i]
		_, err := client.CreateUser(user).Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func (fs *FullstackSuiteQuery) TearDownSuite() {

	for i := 0; i < len(fs.queryUsers); i++ {
		user := fs.queryUsers[i]
		_, err := client.DeleteUser(prisma.UserWhereUniqueInput{
			Email: &user.Email,
		}).Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func (fs *FullstackSuiteQuery) TestQueryUsers() {

}

func (fs *FullstackSuiteQuery) TestQueryPosts() {

}

func TestQuery(t *testing.T) {
	suite.Run(t, new(FullstackSuiteQuery))
}
