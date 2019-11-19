package fullstack_backend

import (
	// "log"
	"testing"

	// models "github.com/german9304/fullstack-backend/models"
	// prisma "github.com/german9304/fullstack-backend/prisma-client"
	// "github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

type FullstackSuiteQuery struct {
	suite.Suite
}


func (fs *FullstackSuiteQuery) SetupSuite() {
	

	
}


func (fs *FullstackSuiteQuery) TearDownSuite() {



}

// func (fs *FullstackSuiteQuery) TestQueryUsers() {
// 	USERS := `
// 		query Users{
// 			users {
// 				id
// 				email
// 				name
// 				password
// 				posts {
// 					id
// 					text
// 				}
// 				likes {
// 					id
// 					quantity
// 				}
// 			}
// 		}
// 	`

// 	usersReq := graphql.NewRequest(USERS)

// 	usersReq.Header.Set("Cache-Control", "no-cache")

// 	var usersRespData map[string][]models.User
// 	if err := clientGraphql.Run(ctx, usersReq, &usersRespData); err != nil {
// 		log.Fatal(err)
// 	}

// 	users := usersRespData["users"]
// 	for _, v := range users {
// 		fs.Assert().NotEmpty(v.Id)
// 		fs.Assert().Len(v.Posts, 1)
// 		fs.Assert().Len(v.Likes, 1)
// 		fs.Assert().NotEmpty(v.Posts[0].Text)
// 		fs.Assert().Equal(int32(0), *v.Likes[0].Quantity)
// 	}
// }

// func (fs *FullstackSuiteQuery) TestQueryPosts() {
// 	POSTS := `
// 		query queryPosts {
// 			posts {
// 				id
// 				text
// 				author {
// 					email
// 					name
// 				}
// 				likes {
// 					quantity
// 				}
// 			}
// 		}
// 	`
// 	postsReq := graphql.NewRequest(POSTS)

// 	postsReq.Header.Set("Cache-Control", "no-cache")

// 	var postsRespData map[string][]models.Post
// 	if err := clientGraphql.Run(ctx, postsReq, &postsRespData); err != nil {
// 		log.Printf("error: %v \n", err)
// 		log.Fatal(err)
// 	}

// 	posts := postsRespData["posts"]
// 	for _, v := range posts {
// 		fs.Assert().NotEmpty(v.Id)
// 		fs.Assert().Len(v.Likes, 1)
// 		fs.Assert().NotEmpty(v.Author)
// 		fs.Assert().Equal(int32(0), *v.Likes[0].Quantity)
// 	}
// }

// func (fs *FullstackSuiteQuery) TestQueryLikes() {
// 	LIKES := `
// 		query queryLikes {
// 			likes {
// 				id
// 				quantity
// 				user {
// 					id
// 				}
// 				post {
// 					id
// 				}
// 			}
// 		}
// 	`
// 	likesReq := graphql.NewRequest(LIKES)

// 	likesReq.Header.Set("Cache-Control", "no-cache")

// 	var likesRespData map[string][]models.Like
// 	if err := clientGraphql.Run(ctx, likesReq, &likesRespData); err != nil {
// 		log.Printf("error: %v \n", err)
// 		log.Fatal(err)
// 	}

// 	likes := likesRespData["likes"]
// 	for _, v := range likes {
// 		fs.Assert().NotEmpty(v.Id)
// 		fs.Assert().Equal(0, v.Quantity)
// 		fs.Assert().NotEmpty(v.User.ID)
// 		fs.Assert().NotEmpty(v.Post.ID)
// 	}
// }


// func (fs *FullstackSuiteQuery) TestQuerySingle() {
// 	USERBYID := `
// 		query UserID($id: String!) {
// 			userById(id: $id) {
// 				id
// 				name
// 				email
// 				posts {
// 					id
// 					text
// 				}
// 				likes {
// 					id
// 					quantity
// 				}
// 			}
// 		}
// 	`

// 	USERBYEMAIL := `
// 		query UserEmail($email: String!) {
// 			userByEmail(email: $email) {
// 				id
// 				name
// 				email
// 				posts {
// 					id
// 					text
// 				}
// 				likes {
// 					id
// 					quantity
// 				}
// 			}
// 		}
// 	`

// 	POST := `
// 		query postQuery($id: String!) {
// 			post(id: $id) {
// 				id
// 				text
// 				author {
// 					id
// 					email
// 				}
// 				likes {
// 					id
// 					quantity
// 				}
// 			}
// 		}
// 	`

// 	LIKE := `
// 		query likeQuery($id: String!) {
// 		    like(id: $id) {
// 				id
// 				quantity
// 				user {
// 					id
// 					email
// 				}
// 				post {
// 					id
// 					text
// 				}
// 			}
// 		}
// 	`

// 	userByIdReq := clientReq(USERBYID, fs.userIds[0], "id")
// 	userByEmailReq := clientReq(USERBYEMAIL, "pepe@mail.com", "email")

// 	var userRespDataById map[string]models.User
// 	if err := clientGraphql.Run(ctx, userByIdReq, &userRespDataById); err != nil {
// 		log.Fatal(err)
// 	}

// 	var userRespDataByEmail map[string]models.User
// 	if err := clientGraphql.Run(ctx, userByEmailReq, &userRespDataByEmail); err != nil {
// 		log.Fatal(err)
// 	}

// 	userById := userRespDataById["userById"]
// 	userByEmail := userRespDataByEmail["userByEmail"]

// 	userTest := func(user models.User) {
// 		fs.Assert().NotEmpty(user.Id)
// 		fs.Assert().Equal("pepe", user.Name)
// 		fs.Assert().Equal("pepe@mail.com", user.Email)
// 		fs.Assert().Equal(1, len(user.Posts))
// 		fs.Assert().Equal(1, len(user.Likes))
// 		fs.Assert().Equal("post1", user.Posts[0].Text)
// 	}

// 	userTest(userById)
// 	userTest(userByEmail)

// 	postId := userById.Posts[0].ID
// 	likeId := userById.Likes[0].ID
// 	postReq := clientReq(POST, postId, "id")
// 	likeReq := clientReq(LIKE, likeId, "id")

// 	var postReqData map[string]models.Post
// 	if err := clientGraphql.Run(ctx, postReq, &postReqData); err != nil {
// 		log.Fatal(err)
// 	}

// 	var likeReqData map[string]models.Like
// 	if err := clientGraphql.Run(ctx, likeReq, &likeReqData); err != nil {
// 		log.Fatal(err)
// 	}

// 	postById := postReqData["post"]
// 	likeById := likeReqData["like"]


// 	fs.Assert().NotEmpty(postById.Id)
// 	fs.Assert().NotEmpty(likeById.Id)
// 	fs.Assert().Equal("pepe@mail.com", likeById.User.Email)
// 	fs.Assert().Equal("pepe@mail.com", postById.Author.Email)
// 	fs.Assert().Equal(int32(0), int32(*postById.Likes[0].Quantity))
// 	fs.Assert().Equal(0, likeById.Quantity)

// }

func TestQuery(t *testing.T) {
	suite.Run(t, new(FullstackSuiteQuery))
}
