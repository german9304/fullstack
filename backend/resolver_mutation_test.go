package fullstack_backend

import (
	"context"
	"log"
	"testing"
	"time"

	// "github.com/stretchr/testify/assert"
	prisma "github.com/german9304/fullstack-backend/prisma-client"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

type FullstackSuiteMutation struct {
	suite.Suite
	postID string
	usrID  string
}

var (
	// client        *prisma.Client  = prisma.New(nil)
	ctx           context.Context = context.TODO()
	email         string          = "John@mail.com"
	clientGraphql *graphql.Client = graphql.NewClient("http://localhost:8000/")
)

func (fs *FullstackSuiteMutation) BeforeTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)
	name := "John"
	password := "293902122"

	usr, _ := client.CreateUser(prisma.UserCreateInput{
		Email:    email,
		Name:     name,
		Password: password,
	}).Exec(ctx)

	post, _ := client.CreatePost(prisma.PostCreateInput{
		Text: "testpost",
		Author: &prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &usr.ID,
			},
		},
	}).Exec(ctx)

	fs.usrID = usr.ID
	fs.postID = post.ID
}

func (fs *FullstackSuiteMutation) AfterTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)
	userEmail := "mark@mail.com"
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &userEmail,
	}).Exec(ctx)
	client.DeletePost(prisma.PostWhereUniqueInput{
		ID: &fs.postID,
	}).Exec(ctx)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (fs *FullstackSuiteMutation) TestMutationCreate() {

	CREATE_USER := `
		mutation signupMutation($userinput: UserInput!) {
			signup (usrinpt: $userinput) {
				id
				email
				name
				password
			}
		}
	`

	CREATE_POST := `
		mutation postMutation($postinput: PostInput!) {
			createPost (pstinpt: $postinput) {
				id
				text
				createdAt
				updatedAt
				author {
					id
					email
					password
				}
				likes {
					createdAt
				}
			}
		}
	`

	// Create a user
	signupReq := graphql.NewRequest(CREATE_USER)

	usr := UserInput{"mark@mail.com", "Mark", "2923ij3j3"}

	signupReq.Var("userinput", usr)

	signupReq.Header.Set("Cache-Control", "no-cache")

	// run it and capture the response
	var newUserRespData map[string]prisma.User
	if err := clientGraphql.Run(ctx, signupReq, &newUserRespData); err != nil {
		log.Fatal(err)
	}
	newUser := newUserRespData["signup"]

	newUserId := newUser.ID

	fs.Assert().Equal(usr.Email, newUser.Email)
	fs.Assert().Equal(usr.Name, newUser.Name)
	fs.Assert().Equal(usr.Password, newUser.Password)

	// Create a post
	newPostReq := graphql.NewRequest(CREATE_POST)

	post := PostInput{"first", newUserId}

	newPostReq.Var("postinput", post)

	newPostReq.Header.Set("Cache-Control", "no-cache")

	// run it and capture the response

	type PostWithAuthor struct {
		Id        string
		Text      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Author    prisma.User
		Likes     []prisma.Likes
	}
	var newPostRespData map[string]PostWithAuthor
	if err := clientGraphql.Run(ctx, newPostReq, &newPostRespData); err != nil {
		log.Fatal(err)
	}

	// var postD map[string]prisma.Post
	requestedPost := newPostRespData["createPost"]
	// postId := requestedPost.Id
	postText := requestedPost.Text
	authorPost := requestedPost.Author
	authorLikes := requestedPost.Likes

	// testing author fields
	fs.Assert().Equal(post.Text, postText)
	fs.Assert().Equal(usr.Email, authorPost.Email)
	fs.Assert().Equal(0, len(authorLikes))

}

func (fs *FullstackSuiteMutation) TestMutationUpdates() {
	UPDATE_POST := `
		mutation updatePostMutation($id: String, $text: String!) {
			updatePost(id: $id, text: $text) {
				id
				text
			}
		}
	
	`
	// Create a user
	updatePostReq := graphql.NewRequest(UPDATE_POST)
	postText := "edited post"
	updatePostReq.Var("id", fs.postID)
	updatePostReq.Var("text", postText)

	updatePostReq.Header.Set("Cache-Control", "no-cache")

	// run it and capture the response
	var newUpdatePostRespData map[string]prisma.Post
	if err := clientGraphql.Run(ctx, updatePostReq, &newUpdatePostRespData); err != nil {
		log.Fatal(err)
	}

	updatedPost := newUpdatePostRespData["updatePost"]
	fs.Assert().Equal(updatedPost.Text, postText)
}

func (fs *FullstackSuiteMutation) TestMutationDelete() {

	post2, _ := client.CreatePost(prisma.PostCreateInput{
		Text: "test2post",
		Author: &prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &fs.usrID,
			},
		},
	}).Exec(ctx)

	DELETE_POST := `
		mutation deletePostMutation($id: String){
			deletePost(id: $id) {
				text
			}
		}
	`

	deletePostReq := graphql.NewRequest(DELETE_POST)
	deletePostReq.Var("id", post2.ID)

	deletePostReq.Header.Set("Cache-Control", "no-cache")

	// run it and capture the response
	var deletePostRespData map[string]prisma.Post
	if err := clientGraphql.Run(ctx, deletePostReq, &deletePostRespData); err != nil {
		log.Printf("error: %v \n", err)
		log.Fatal(err)
	}

	p, err := client.Post(prisma.PostWhereUniqueInput{
		ID: &post2.ID,
	}).Exec(ctx)

	deletePost := deletePostRespData["deletePost"]
	fs.Assert().Equal(deletePost.Text, "test2post")
	fs.Assert().Equal("query returned no result", err.Error())
	fs.Assert().Nil(p)
}

func TestSetSuite(t *testing.T) {
	suite.Run(t, new(FullstackSuiteMutation))
}
