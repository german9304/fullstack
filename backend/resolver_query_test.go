package fullstack_backend

import (
	"log"
	"testing"

	models "github.com/german9304/fullstack-backend/models"
	prisma "github.com/german9304/fullstack-backend/prisma-client"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

type FullstackSuiteQuery struct {
	suite.Suite
	queryUsers    []prisma.UserCreateInput
	clientGraphql *graphql.Client
}

var (
	testingModelUsers []models.User = []models.User{
		models.User{
			Email:    "pepe@mail.com",
			Name:     "pepe",
			Password: "pepe1234",
		},
		models.User{
			Email:    "jimmy@mail.com",
			Name:     "Jimmy",
			Password: "jim1234",
		},
		models.User{
			Email:    "miguel@mail.com",
			Name:     "Miguel",
			Password: "miguel1234",
		},
		models.User{
			Email:    "kendrick@mail.com",
			Name:     "Kendrick",
			Password: "kendrick1234",
		},
	}
	testingPostsNames []string = []string{
		"post1",
		"post2",
		"post3",
		"post4",
	}
)

func (fs *FullstackSuiteQuery) SetupSuite() {

	fs.clientGraphql = graphql.NewClient("http://localhost:8000/")

	for i := 0; i < len(testingModelUsers); i++ {
		user := testingModelUsers[i]
		userInput := prisma.UserCreateInput{
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
			Posts: &prisma.PostCreateManyWithoutAuthorInput{
				Create: []prisma.PostCreateWithoutAuthorInput{
					prisma.PostCreateWithoutAuthorInput{
						Text: testingPostsNames[i],
						Likes: &prisma.LikesCreateManyWithoutPostInput{
							Create: []prisma.LikesCreateWithoutPostInput{
								prisma.LikesCreateWithoutPostInput{
									User: prisma.UserCreateOneWithoutLikesInput{
										Connect: &prisma.UserWhereUniqueInput{
											Email: &user.Email,
										},
									},
								},
							},
						},
					},
				},
			},
		}

		fs.queryUsers = append(fs.queryUsers, userInput)
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
	USERS := `
		query Users{
			users {
				id
				email
				name
				password
			}
		}
	`
	usersReq := graphql.NewRequest(USERS)

	usersReq.Header.Set("Cache-Control", "no-cache")

	var usersRespData map[string][]models.User
	if err := clientGraphql.Run(ctx, usersReq, &usersRespData); err != nil {
		log.Printf("error: %v \n", err)
		log.Fatal(err)
	}

	users := usersRespData["users"]
	for k, v := range users {
		fs.Assert().NotEmpty(v.Id)
		fs.Assert().Equal(testingModelUsers[k].Email, v.Email)
		fs.Assert().Equal(testingModelUsers[k].Name, v.Name)
		fs.Assert().Equal(testingModelUsers[k].Password, v.Password)
	}
}

func (fs *FullstackSuiteQuery) TestQueryPosts() {

}

func TestQuery(t *testing.T) {
	suite.Run(t, new(FullstackSuiteQuery))
}
