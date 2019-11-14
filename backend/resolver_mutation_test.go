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
	postID        string
	usrID         string
	commentID     string
	userEmail     string
	userPassword  string
	commentLikeID string
	postLikeID    string
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

	postLike, _ := client.CreateLikePost(prisma.LikePostCreateInput{
		User: prisma.UserCreateOneWithoutPostLikesInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
		Post: prisma.PostCreateOneWithoutLikesInput{
			Connect: &prisma.PostWhereUniqueInput{
				ID: &post.ID,
			},
		},
	}).Exec(ctx)

	commentLike, _ := client.CreateLikeComment(prisma.LikeCommentCreateInput{
		User: prisma.UserCreateOneWithoutCommentLikesInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
		Comment: prisma.CommentCreateOneWithoutLikesInput{
			Connect: &prisma.CommentWhereUniqueInput{
				ID: &comment.ID,
			},
		},
	}).Exec(ctx)

	fs.usrID = user.ID
	fs.postID = post.ID
	fs.commentID = comment.ID
	fs.commentLikeID = commentLike.ID
	fs.postLikeID = postLike.ID
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

	const CREATEPOSTLIKE string = `
		mutation createLikeFromPost($likeinput: PostLikeInput!) {
			createPostLike(likeinput: $likeinput) {
				id
				quantity
			}
		}
	`

	const CREATECOMMENTLIKE string = `
		mutation createLikeFromPost($likeinput: CommentLikeInput!) {
			createCommentLike(likeinput: $likeinput) {
				id
				quantity
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

	likeInputParam := LikeInput{0, fs.usrID}

	postLikeparam := PostLikeInput{&likeInputParam, fs.postID}
	createPostLikeParams := []RequestParams{
		RequestParams{"likeinput", postLikeparam},
	}
	commentLikeparam := CommentLikeInput{&likeInputParam, fs.commentID}
	createCommentLikeParams := []RequestParams{
		RequestParams{"likeinput", commentLikeparam},
	}

	createPostLikeReq := clientRequests(CREATEPOSTLIKE, createPostLikeParams)
	createCommentLikeReq := clientRequests(CREATECOMMENTLIKE, createCommentLikeParams)
	createdPostLike := createPostLikeReq["createPostLike"].(map[string]interface{})
	createdCommentLike := createCommentLikeReq["createCommentLike"].(map[string]interface{})

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
	//Likes
	fs.Assert().NotEmpty(createdPostLike["id"].(string))
	fs.Assert().NotEmpty(createdCommentLike["id"].(string))
	fs.Assert().Equal(float64(0), createdCommentLike["quantity"].(float64))
	fs.Assert().Equal(float64(0), createdPostLike["quantity"].(float64))
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

	const UPDATELIKECOMMENT string = `
		mutation updateLikeComment($id: String!, $quantity: Int!, $liketype: String!) {
			updateLike(id: $id, quantity: $quantity, liketype: $liketype) {
				id
				quantity
			}
		}
	`

	const UPDATELIKEPOST string = `
		mutation updateLikePost($id: String!, $quantity: Int!, $liketype: String!) {
			updateLike(id: $id, quantity: $quantity, liketype: $liketype) {
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

	updatelikeParams := func(quantity int, id, liketype string) []RequestParams {
		params := []RequestParams{
			RequestParams{"id", id},
			RequestParams{"quantity", quantity},
			RequestParams{"liketype", liketype},
		}
		return params
	}

	updateLikeCommentParams := updatelikeParams(2, fs.commentLikeID, "comment")
	updateLikePostParams := updatelikeParams(3, fs.postLikeID, "post")

	updatedLikeCommentReq := clientRequests(UPDATELIKECOMMENT, updateLikeCommentParams)
	updatedLikePostReq := clientRequests(UPDATELIKECOMMENT, updateLikePostParams)

	updatedLikeComment := updatedLikeCommentReq["updateLike"].(map[string]interface{})
	updatedLikePost := updatedLikePostReq["updateLike"].(map[string]interface{})

	// Updated Post
	fs.Assert().NotEmpty(updatedPost["id"].(string))
	fs.Assert().Equal(updatePostInput.Body, updatedPost["body"].(string))
	fs.Assert().Equal(updatePostInput.Header, updatedPost["header"].(string))
	// Updated comment
	fs.Assert().NotEmpty(updatedComment["id"].(string))
	fs.Assert().Equal(commentBody, updatedComment["body"].(string))
	//Updated Likes
	fs.Assert().Equal(float64(2), updatedLikeComment["quantity"].(float64))
	fs.Assert().Equal(float64(3), updatedLikePost["quantity"].(float64))
}

func (fs *FullstackSuiteMutation) TestMutationDelete() {

	post2, _ := client.CreatePost(prisma.PostCreateInput{
		Header:  "header from post2",
		Picture: nil,
		Body:    "body from post2",
		Author: prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &fs.usrID,
			},
		},
	}).Exec(ctx)

	comment2, _ := client.CreateComment(prisma.CommentCreateInput{
		Body:    "body from comment2",
		Author: prisma.UserCreateOneWithoutCommentsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &fs.usrID,
			},
		},
		Post: prisma.PostCreateOneWithoutCommentsInput{
			Connect: &prisma.PostWhereUniqueInput{
				ID: &post2.ID,
			},
		},
	}).Exec(ctx)

	const DELETEPOST string = `
		mutation deletePostMutation($id: String!){
			deletePost(id: $id) {
				id
				header
			}
		}
	`
	const DELETECOMMENT string = `
		mutation deleteCommentMutation($id: String!){
			deleteComment(id: $id) {
				id
				body
			}
		}
	`

	deletePostParams := []RequestParams{
		RequestParams{"id", post2.ID},
	}

	deleteCommentParams := []RequestParams{
		RequestParams{"id", comment2.ID},
	}

	deleteCommentReq := clientRequests(DELETECOMMENT, deleteCommentParams)
	deletedComment := deleteCommentReq["deleteComment"].(map[string]interface{})
	deletedCommentId := deletedComment["id"].(string)
	_, deletedCommenterr := client.Comment(prisma.CommentWhereUniqueInput{
		ID: &deletedCommentId,
	}).Exec(ctx)

	fs.EqualError(deletedCommenterr, "query returned no result")

	deletePostReq := clientRequests(DELETEPOST, deletePostParams)
	deletedPost := deletePostReq["deletePost"].(map[string]interface{})
	deletedPostId := deletedPost["id"].(string)
	_, deletedPosterr := client.Post(prisma.PostWhereUniqueInput{
		ID: &deletedPostId,
	}).Exec(ctx)


	fs.EqualError(deletedPosterr, "query returned no result")

}

func TestMutaion(t *testing.T) {
	suite.Run(t, new(FullstackSuiteMutation))
}
