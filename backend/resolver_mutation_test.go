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
	commentID    string
	userEmail    string
	userPassword string
}

func (fs *FullstackSuiteMutation) SetupSuite() {
	name := "John"
	fs.userPassword = "293902122"
	fs.userEmail = "John@mail.com"

	user, _ := client.CreateUser(prisma.UserCreateInput{
		Email:    fs.userEmail,
		Name:     name,
		Password: fs.userPassword,
	}).Exec(ctx)

	post, _ := client.CreatePost(prisma.PostCreateInput{
		Body:    "testbody",
		Header:  "testpost",
		Picture: nil,
		Author: prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
	}).Exec(ctx)

	comment, _ := client.CreateComment(prisma.CommentCreateInput{
		Body: "commentBody",
		Author: prisma.UserCreateOneWithoutCommentsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
		Post: prisma.PostCreateOneWithoutCommentsInput{
			Connect: &prisma.PostWhereUniqueInput{
				ID: &post.ID,
			},
		},
	},
	).Exec(ctx)

	fs.usrID = user.ID
	fs.postID = post.ID
	fs.commentID = comment.ID
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

	userInput := UserInput{"mark@mail.com", "Mark", "2923ij3j3"}
	signUpParams := []RequestParams{
		RequestParams{"userinput", userInput},
	}
	signupReq := clientRequests(CREATEUSER, signUpParams)
	newUser := signupReq["signup"].(map[string]interface{})

	newUserId := newUser["id"].(string)
	postInput := PostInput{newUserId, "header1", "body1"}

	createPostParams := []RequestParams{
		RequestParams{"postinput", postInput},
		RequestParams{"pic", nil},
	}

	createPostReq := clientRequests(CREATEPOST, createPostParams)
	newPost := createPostReq["createPost"].(map[string]interface{})
	newPostId := newPost["id"].(string)

	commentInput := CommentInput{"Comment body", newUserId, newPostId}
	createCommentParams := []RequestParams{
		RequestParams{"commentinput", commentInput},
	}

	createCommentReq := clientRequests(CREATECOMMENT, createCommentParams)
	newComment := createCommentReq["createComment"].(map[string]interface{})

	// User testing
	fs.Assert().NotEmpty(newUserId)
	fs.Assert().Equal(userInput.Email, newUser["email"].(string))
	fs.Assert().Equal(userInput.Name, newUser["name"].(string))
	fs.Assert().Equal(userInput.Password, newUser["password"].(string))
	// Post testing
	fs.Assert().NotEmpty(newPostId)
	fs.Assert().Equal(postInput.Header, newPost["header"].(string))
	fs.Assert().Equal(postInput.Body, newPost["body"].(string))
	// Comment testing
	fs.Assert().NotEmpty(newComment["id"].(string))
	fs.Assert().Equal(commentInput.Body, newComment["body"].(string))
}

func (fs *FullstackSuiteMutation) TestMutationUpdates() {
	const UPDATEPOST string = `
		mutation updatePostMutation($id: String!, $postinput: UpdatePostInput!, $pic: Upload) {
			updatePost(id: $id, postinput: $postinput, picture: $pic) {
				id
				header
				body
			}
		}

	`
	const UPDATECOMMENT string = `
		mutation updateCommentMutation($id: String!, $body: String!) {
			updateComment(id: $id, body: $body) {
				id
				body
			}
		}
	`

	const UPDATELIKE string = `
		mutation updateLikeFromPost($id: String!, $quantity: Int!) {
			updateLike(id: $id, quantity: $quantity) {
				id
				quantity
			}
		}
	`

	updatePostInput := UpdatePostInput{"updatedHeader", "updatedBody"}
	updatePostParams := []RequestParams{
		RequestParams{"id", fs.postID},
		RequestParams{"postinput", updatePostInput},
		RequestParams{"pic", nil},
	}

	updatePostReq := clientRequests(UPDATEPOST, updatePostParams)
	updatedPost := updatePostReq["updatePost"].(map[string]interface{})

	commentBody := "newbody"
	updateCommentParams := []RequestParams{
		RequestParams{"id", fs.commentID},
		RequestParams{"body", commentBody},
	}

	updateCommentReq := clientRequests(UPDATECOMMENT, updateCommentParams)
	updatedComment := updateCommentReq["updateComment"].(map[string]interface{})

	likeInputParam := LikeInput{fs.usrID, 1}
	postLikeparam := PostLikeInput{&likeInputParam, fs.postID}
	updatePostLikeParams := []RequestParams{
		RequestParams{"likeinput", postLikeparam},
	}
	updateLikeReq := clientRequests(UPDATELIKE, updatePostLikeParams)

	log.Printf("like result => %v \n", updateLikeReq)
	// Updated Post
	fs.Assert().NotEmpty(updatedPost["id"].(string))
	fs.Assert().Equal(updatePostInput.Body, updatedPost["body"].(string))
	fs.Assert().Equal(updatePostInput.Header, updatedPost["header"].(string))
	// Updated comment
	fs.Assert().NotEmpty(updatedComment["id"].(string))
	fs.Assert().Equal(commentBody, updatedComment["body"].(string))
}

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
