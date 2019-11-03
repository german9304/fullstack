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
				posts {
					id
					text
				}
				likes {
					id
					quantity
				}
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
	for _, v := range users {
		fs.Assert().NotEmpty(v.Id)
		fs.Assert().Len(v.Posts, 1)
		fs.Assert().Len(v.Likes, 1)
		fs.Assert().NotEmpty(v.Posts[0].Text)
		fs.Assert().Equal(int32(0), *v.Likes[0].Quantity)
	}
}

func (fs *FullstackSuiteQuery) TestQueryPosts() {
	POSTS := `
		query queryPosts {
			posts {
				id
				text
				author {
					email
					name
				}
				likes {
					quantity
				}
			}
		}
	`
	postsReq := graphql.NewRequest(POSTS)

	postsReq.Header.Set("Cache-Control", "no-cache")

	var postsRespData map[string][]models.Post
	if err := clientGraphql.Run(ctx, postsReq, &postsRespData); err != nil {
		log.Printf("error: %v \n", err)
		log.Fatal(err)
	}

	posts := postsRespData["posts"]
	for _, v := range posts {
		fs.Assert().NotEmpty(v.Id)
		fs.Assert().Len(v.Likes, 1)
		fs.Assert().NotEmpty(v.Author)
		fs.Assert().Equal(int32(0), *v.Likes[0].Quantity)
	}
}

func (fs *FullstackSuiteQuery) TestQueryLikes() {
	LIKES := `
		query queryLikes {
			likes {
				id
				quantity
			}
		}
	`
	likesReq := graphql.NewRequest(LIKES)

	likesReq.Header.Set("Cache-Control", "no-cache")

	var likesRespData map[string][]models.Like
	if err := clientGraphql.Run(ctx, likesReq, &likesRespData); err != nil {
		log.Printf("error: %v \n", err)
		log.Fatal(err)
	}

	likes := likesRespData["likes"]
	for _, v := range likes {
		fs.Assert().NotEmpty(v.Id)
		log.Printf("quantity %v \n", v.Quantity)
	}
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(FullstackSuiteQuery))
}
