package fullstack_backend

import (
	"context"
	"log"
	"testing"

	// "github.com/stretchr/testify/assert"
	prisma "github.com/german9304/fullstack-backend/prisma-client"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

type FullStackSuite struct {
	suite.Suite
	postID string
}

var (
	// client        *prisma.Client  = prisma.New(nil)
	ctx           context.Context = context.TODO()
	email         string          = "John@mail.com"
	clientGraphql *graphql.Client = graphql.NewClient("http://localhost:8000/")
)

func (fs *FullStackSuite) BeforeTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)
	name := "John"
	password := "293902122"

	client.CreateUser(prisma.UserCreateInput{
		Email:    email,
		Name:     name,
		Password: password,
	}).Exec(ctx)
}

func (fs *FullStackSuite) AfterTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)
	userEmail := "mark@mail.com"
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &userEmail,
	}).Exec(ctx)
	log.Printf("post id => %v \n", fs.postID)
	client.DeletePost(prisma.PostWhereUniqueInput{
		ID: &fs.postID,
	}).Exec(ctx)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (fs *FullStackSuite) TestMutations() {

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

	ctx := context.Background()

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
		Id     string
		Text   string
		Author prisma.User
		Likes  []prisma.Likes
	}
	var newPostRespData map[string]PostWithAuthor
	if err := clientGraphql.Run(ctx, newPostReq, &newPostRespData); err != nil {
		log.Fatal(err)
	}

	// var postD map[string]prisma.Post
	requestedPost := newPostRespData["createPost"]
	postId := requestedPost.Id
	postText := requestedPost.Text
	authorPost := requestedPost.Author
	authorLikes := requestedPost.Likes

	fs.postID = postId

	log.Printf("id: %v, text: %v \n", postId, postText)
	log.Printf("Author: %v \n", authorPost)
	log.Printf("Likes: %v \n", authorLikes)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSetSuite(t *testing.T) {
	suite.Run(t, new(FullStackSuite))
}
