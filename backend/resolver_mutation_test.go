package fullstack_backend

import (
	"log"
	"testing"

	// "github.com/stretchr/testify/assert"
	// gqlgengraphql "github.com/99designs/gqlgen/graphql"
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

	client.CreateUser(prisma.UserCreateInput{
		Email:    fs.userEmail,
		Name:     name,
		Password: fs.userPassword,
	}).Exec(ctx)

	// post, _ := client.CreatePost(prisma.PostCreateInput{
	// 	Text: "testpost",
	// 	Author: &prisma.UserCreateOneWithoutPostsInput{
	// 		Connect: &prisma.UserWhereUniqueInput{
	// 			ID: &usr.ID,
	// 		},
	// 	},
	// }).Exec(ctx)

	// fs.usrID = usr.ID
	// fs.postID = post.ID
}

func (fs *FullstackSuiteMutation) TearDownSuite() {
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &fs.userEmail,
	}).Exec(ctx)
	userEmail := "mark@mail.com"
	client.DeleteUser(prisma.UserWhereUniqueInput{
		Email: &userEmail,
	}).Exec(ctx)
	// client.DeletePost(prisma.PostWhereUniqueInput{
	// 	ID: &fs.postID,
	// }).Exec(ctx)
}

type RequestParams struct {
	key   string
	value interface{}
}


func clientRequests(request string, reqParams []RequestParams) map[string]interface{} {
	clientRequest := graphql.NewRequest(request)
	for _, v := range reqParams {
		clientRequest.Var(v.key, v.value)
	}
	clientRequest.Header.Set("Cache-Control", "no-cache")
	var newRespData map[string]interface{}
	if err := clientGraphql.Run(ctx, clientRequest, &newRespData); err != nil {
		log.Fatal(err)
	}
	return newRespData
}

func (fs *FullstackSuiteMutation) TestMutationCreate() {

	const CREATEUSER string = `
		mutation signupMutation($userinput: UserInput!) {
			signup (usrinpt: $userinput) {
				id
				email
				name
				password
			}
		}
	`

	const CREATEPOST string = `
		mutation postMutation($postinput: PostInput!, $pic: Upload) {
			createPost (pstinpt: $postinput, picture: $pic) {
				id
				body
				header
			}
		}
	`

	const CREATECOMMENT string = `
		mutation createCommentMutation($commentinput: CommentInput!) {
			createComment(commentinput: $commentinput) {
				id
				body
			}
		}
	`

	const CREATELIKE string = `
		mutation likeMutation($likeinput: LikeInput!) {
			createLike(likeInput: $likeinput) {
				id
				quantity
			}
		}
	`

	signUpParams := []RequestParams{
		RequestParams{"userinput", UserInput{"mark@mail.com", "Mark", "2923ij3j3"}},
	}
	signupReq := clientRequests(CREATEUSER, signUpParams)

	newUser := signupReq["signup"].(map[string]interface{})

	log.Printf("%v \n", newUser)

	newUserId := newUser["id"].(string)

	createPostParams := []RequestParams{
		RequestParams{"postinput", PostInput{newUserId, "header1", "body1"}},
		RequestParams{"pic", nil},
	}

	createPostReq := clientRequests(CREATEPOST, createPostParams)

	newPost := createPostReq["createPost"].(map[string]interface{})

	log.Printf("Created post %v \n", newPost)

	newPostId := newPost["id"].(string)

	createCommentParams := []RequestParams{
		RequestParams{"commentinput", CommentInput{"This is a paragraph", newUserId, newPostId}},
	}
	
	createCommentReq := clientRequests(CREATECOMMENT, createCommentParams)

	newComment := createCommentReq["createComment"].(map[string]interface{})

	log.Printf("Created comment %v \n", newComment)

	// likeInput := LikeInput{newUserId, postId, 1}

	// likesReq.Var("likeinput", likeInput)

	// likesReq.Header.Set("Cache-Control", "no-cache")

	// var newLikesRespData map[string]prisma.Like
	// if err := clientGraphql.Run(ctx, likesReq, &newLikesRespData); err != nil {
	// 	log.Fatal(err)
	// }

	// likesResp := newLikesRespData["createLike"]
	// likesQuantity := *likesResp.Quantity

	// fs.Assert().Equal(usr.Email, newUser.Email)
	// fs.Assert().Equal(usr.Name, newUser.Name)
	// fs.Assert().Equal(usr.Password, newUser.Password)
	// fs.Assert().Equal(post.Text, postText)
	// fs.Assert().Equal(int32(1), likesQuantity)
}

// func (fs *FullstackSuiteMutation) TestMutationUpdates() {
// 	UPDATE_POST := `
// 		mutation updatePostMutation($id: String!, $text: String!) {
// 			updatePost(id: $id, text: $text) {
// 				id
// 				text
// 			}
// 		}

// 	`
// 	// Create a user
// 	updatePostReq := graphql.NewRequest(UPDATE_POST)
// 	postText := "edited post"
// 	updatePostReq.Var("id", fs.postID)
// 	updatePostReq.Var("text", postText)

// 	updatePostReq.Header.Set("Cache-Control", "no-cache")

// 	// run it and capture the response
// 	var newUpdatePostRespData map[string]prisma.Post
// 	if err := clientGraphql.Run(ctx, updatePostReq, &newUpdatePostRespData); err != nil {
// 		log.Fatal(err)
// 	}

// 	updatedPost := newUpdatePostRespData["updatePost"]
// 	fs.Assert().Equal(updatedPost.Text, postText)
// }

// func (fs *FullstackSuiteMutation) TestMutationDelete() {

// 	post2, _ := client.CreatePost(prisma.PostCreateInput{
// 		Text: "test2post",
// 		Author: &prisma.UserCreateOneWithoutPostsInput{
// 			Connect: &prisma.UserWhereUniqueInput{
// 				ID: &fs.usrID,
// 			},
// 		},
// 	}).Exec(ctx)

// 	DELETE_POST := `
// 		mutation deletePostMutation($id: String!){
// 			deletePost(id: $id) {
// 				text
// 			}
// 		}
// 	`

// 	deletePostReq := graphql.NewRequest(DELETE_POST)
// 	deletePostReq.Var("id", post2.ID)

// 	deletePostReq.Header.Set("Cache-Control", "no-cache")

// 	// run it and capture the response
// 	var deletePostRespData map[string]prisma.Post
// 	if err := clientGraphql.Run(ctx, deletePostReq, &deletePostRespData); err != nil {
// 		log.Printf("error: %v \n", err)
// 		log.Fatal(err)
// 	}

// 	p, err := client.Post(prisma.PostWhereUniqueInput{
// 		ID: &post2.ID,
// 	}).Exec(ctx)

// 	deletePost := deletePostRespData["deletePost"]
// 	fs.Assert().Equal(deletePost.Text, "test2post")
// 	fs.Assert().Equal("query returned no result", err.Error())
// 	fs.Assert().Nil(p)
// }

func TestMutaion(t *testing.T) {
	suite.Run(t, new(FullstackSuiteMutation))
}
