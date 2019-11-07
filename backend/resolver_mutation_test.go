package fullstack_backend

import (
	"log"
	"testing"

	// "github.com/stretchr/testify/assert"
	prisma "github.com/german9304/fullstack-backend/prisma-client"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

type FullstackSuiteMutation struct {
	suite.Suite
	postID       string
	usrID        string
	userEmail    string
	userPassword string
}

func (fs *FullstackSuiteMutation) SetupSuite() {
	name := "John"
	fs.userPassword = "293902122"
	fs.userEmail = "John@mail.com"

	usr, _ := client.CreateUser(prisma.UserCreateInput{
		Email:    fs.userEmail,
		Name:     name,
		Password: fs.userPassword,
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

func (fs *FullstackSuiteMutation) TearDownSuite() {
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &fs.userEmail,
	}).Exec(ctx)
	userEmail := "mark@mail.com"
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &userEmail,
	}).Exec(ctx)
	client.DeletePost(prisma.PostWhereUniqueInput{
		ID: &fs.postID,
	}).Exec(ctx)
}

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
			}
		}
	`

	CREATE_LIKE := `
		mutation likeMutation($likeinput: LikeInput!) {
			createLike(likeInput: $likeinput) {
				id
				quantity
			}
		}
	`

	// Create a user request
	signupReq := graphql.NewRequest(CREATE_USER)
	// Create a post request
	newPostReq := graphql.NewRequest(CREATE_POST)
	// Create a like request
	likesReq := graphql.NewRequest(CREATE_LIKE)

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

	post := PostInput{"first", newUserId}

	newPostReq.Var("postinput", post)

	newPostReq.Header.Set("Cache-Control", "no-cache")

	var newPostRespData map[string]prisma.Post
	if err := clientGraphql.Run(ctx, newPostReq, &newPostRespData); err != nil {
		log.Fatal(err)
	}

	requestedPost := newPostRespData["createPost"]
	postId := requestedPost.ID
	postText := requestedPost.Text

	likeInput := LikeInput{newUserId, postId, 1}

	likesReq.Var("likeinput", likeInput)

	likesReq.Header.Set("Cache-Control", "no-cache")

	var newLikesRespData map[string]prisma.Like
	if err := clientGraphql.Run(ctx, likesReq, &newLikesRespData); err != nil {
		log.Fatal(err)
	}

	likesResp := newLikesRespData["createLike"]
	likesQuantity := *likesResp.Quantity

	fs.Assert().Equal(usr.Email, newUser.Email)
	fs.Assert().Equal(usr.Name, newUser.Name)
	fs.Assert().Equal(usr.Password, newUser.Password)
	fs.Assert().Equal(post.Text, postText)
	fs.Assert().Equal(int32(1), likesQuantity)
}

func (fs *FullstackSuiteMutation) TestMutationUpdates() {
	UPDATE_POST := `
		mutation updatePostMutation($id: String!, $text: String!) {
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
		mutation deletePostMutation($id: String!){
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

func (fs *FullstackSuiteMutation) TestMutationSignIn() {
	const SIGNIN string = `
		mutation userSignin($email: String!, $password: String!) {
			signin(email: $email, password: $password) {
				id
				email
			}
		}
	`

	type reqParams struct {
		reqVar   string
		reqValue string
	}

	clientReq := func(request string, params []reqParams) *graphql.Request {
		clientRequest := graphql.NewRequest(request)
		for _, v := range params {
			clientRequest.Var(v.reqVar, v.reqValue)
		}
		clientRequest.Header.Set("Cache-Control", "no-cache")
		return clientRequest
	}

	paramsWrongPassword := []reqParams{
		reqParams{"email", fs.userEmail},
		reqParams{"password", "user1234"},
	}

	paramsCorrectPassword := []reqParams{
		reqParams{"email", fs.userEmail},
		reqParams{"password", fs.userPassword},
	}

	passwordDoesNotExists := clientReq(SIGNIN, paramsWrongPassword)
	passwordDoesExists := clientReq(SIGNIN, paramsCorrectPassword)

	var signinRespDataInCorrectPswd map[string]prisma.User
	err := clientGraphql.Run(ctx, passwordDoesNotExists, &signinRespDataInCorrectPswd)

	var signinRespDataCorrectPswd map[string]prisma.User
	if err := clientGraphql.Run(ctx, passwordDoesExists, &signinRespDataCorrectPswd); err != nil {
		log.Printf("error: %v \n", err)
		log.Fatal(err)
	}

	fs.Assert().EqualError(err, "graphql: incorrect password, please try again")
	signUser := signinRespDataCorrectPswd["signin"]
	fs.Assert().NotEmpty(signUser.ID)
}

func TestMutaion(t *testing.T) {
	suite.Run(t, new(FullstackSuiteMutation))
}
