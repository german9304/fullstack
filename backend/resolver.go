package fullstack_backend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

var (
	client *prisma.Client = prisma.New(nil)
)

type Resolver struct{}

func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}
func (r *Resolver) LikeCommentResolver() LikeCommentResolver {
	return &LikeCommentResolver{r}
}
func (r *Resolver) LikePostResolver() LikePostResolver {
	return &LikePostResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) CreatedAt(ctx context.Context, obj *prisma.Comment) (*time.Time, error) {
	createdAt := obj.CreatedAt
	t, err := time.Parse(time.RFC3339, createdAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (r *commentResolver) UpdatedAt(ctx context.Context, obj *prisma.Comment) (*time.Time, error) {
	updatedAt := obj.UpdatedAt
	t, err := time.Parse(time.RFC3339, updatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (r *commentResolver) Author(ctx context.Context, obj *prisma.Comment) (*prisma.User, error) {
	commentID := obj.ID
	commentAuthor, err := client.Comment(prisma.CommentWhereUniqueInput{
		ID: &commentID,
	}).Author().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return commentAuthor, nil

}
func (r *commentResolver) Post(ctx context.Context, obj *prisma.Comment) (*prisma.Post, error) {
	commentID := obj.ID
	commentPost, err := client.Comment(prisma.CommentWhereUniqueInput{
		ID: &commentID,
	}).Post().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return commentPost, nil
}
func (r *commentResolver) Likes(ctx context.Context, obj *prisma.Comment) ([]Like, error) {
	panic("not implemented")
	// commentID := obj.ID
	// commentLikes, err := client.Comment(prisma.CommentWhereUniqueInput{
	// 	ID: &commentID,
	// }).Likes(nil).Exec(ctx)

	// if err != nil {
	// 	return nil, err
	// }

	// return commentLikes, nil
}

type LikeCommentResolver struct{ *Resolver }

func (r *LikeCommentResolver) User(ctx context.Context, obj *prisma.LikeComment) (*prisma.User, error) {
	userLikes, err := client.LikeComment(prisma.LikeCommentWhereUniqueInput{
		ID: &obj.ID,
	}).User().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return userLikes, nil
}

func (r *LikeCommentResolver) Comment(ctx context.Context, obj *prisma.LikeComment) (*prisma.Comment, error) {
	commentLike, err := client.LikeComment(prisma.LikeCommentWhereUniqueInput{
		ID: &obj.ID,
	}).Comment().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return commentLike, nil
}
func (r *LikeCommentResolver) CreatedAt(ctx context.Context, obj *prisma.LikeComment) (*time.Time, error) {
	createdAt := obj.CreatedAt
	t, err := time.Parse(time.RFC3339, createdAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (r *LikeCommentResolver) UpdatedAt(ctx context.Context, obj *prisma.LikeComment) (*time.Time, error) {
	UpdatedAt := obj.UpdatedAt
	t, err := time.Parse(time.RFC3339, UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

type LikePostResolver struct{ *Resolver }

func (r *LikeCommentResolver) User(ctx context.Context, obj *prisma.LikePost) (*prisma.User, error) {
	userLikes, err := client.LikePost(prisma.LikePostWhereUniqueInput{
		ID: &obj.ID,
	}).User().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return userLikes, nil
}
func (r *LikePostResolver) Post(ctx context.Context, obj *prisma.LikePost) (*prisma.Post, error) {
	postLike, err := client.LikePost(prisma.LikePostWhereUniqueInput{
		ID: &obj.ID,
	}).Post().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return postLike, nil
}
func (r *LikePostResolver) Comment(ctx context.Context, obj *prisma.LikePost) (*prisma.Comment, error) {
	commentLike, err := client.LikePost(prisma.LikePostWhereUniqueInput{
		ID: &obj.ID,
	}).Comment().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return commentLike, nil
}
func (r *LikePostResolver) CreatedAt(ctx context.Context, obj *prisma.prisma.LikePost) (*time.Time, error) {
	createdAt := obj.CreatedAt
	t, err := time.Parse(time.RFC3339, createdAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (r *LikePostResolver) UpdatedAt(ctx context.Context, obj *prisma.LikePost) (*time.Time, error) {
	UpdatedAt := obj.UpdatedAt
	t, err := time.Parse(time.RFC3339, UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Signup(ctx context.Context, usrinpt UserInput) (*prisma.User, error) {
	email := usrinpt.Email
	name := usrinpt.Name
	pwd := usrinpt.Password

	user, err := client.CreateUser(prisma.UserCreateInput{
		Email:    email,
		Name:     name,
		Password: pwd,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	w := ctx.Value("response").(Auth)
	cookie := http.Cookie{
		Name:     "user",
		Value:    user.ID,
		HttpOnly: true,
	}
	http.SetCookie(w.RW, &cookie)

	return user, nil
}
func (r *mutationResolver) Signin(ctx context.Context, email string, password string) (*prisma.User, error) {
	signedUser, err := client.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	signedUserPassword := signedUser.Password

	if signedUserPassword == password {
		// set cookie to save session
		w := ctx.Value("response").(Auth)
		cookie := http.Cookie{
			Name:     "user",
			Value:    signedUser.ID,
			HttpOnly: true,
		}
		http.SetCookie(w.RW, &cookie)
		return signedUser, nil
	}

	return nil, fmt.Errorf("incorrect password, please try again")
}
func (r *mutationResolver) Signout(ctx context.Context) (*Message, error) {
	w := ctx.Value("response").(Auth)
	cookie, err := w.RQ.Cookie("user")
	if err != nil {
		return nil, err
	}
	ck := http.Cookie{
		Name:     "user",
		Value:    cookie.Value,
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w.RW, &ck)
	message := Message{"Sign out success"}
	return &message, nil
}
func (r *mutationResolver) CreatePost(ctx context.Context, pstinpt PostInput, picture *graphql.Upload) (*prisma.Post, error) {
	author := pstinpt.Author
	header := pstinpt.Header
	pic := picture
	body := pstinpt.Body

	log.Printf("Post infor %v \n", pic)
	newPost := client.CreatePost(prisma.PostCreateInput{
		Header:  header,
		Picture: nil,
		Body:    body,
		Author: prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &author,
			},
		},
	})

	post, err := newPost.Exec(ctx)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, postinput UpdatePostInput, picture *graphql.Upload) (*prisma.Post, error) {
	header := postinput.Header
	// pic := postinput.Picture
	body := postinput.Body
	updatedPost, err := client.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{
			ID: &id,
		},
		Data: prisma.PostUpdateInput{
			Header:  &header,
			Body:    &body,
			Picture: nil,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*prisma.Post, error) {
	deletedPost, err := client.DeletePost(prisma.PostWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return deletedPost, nil
}
func (r *mutationResolver) CreateComment(ctx context.Context, commentinput CommentInput) (*prisma.Comment, error) {
	body := commentinput.Body
	user := commentinput.User
	post := commentinput.Post

	newComment, err := client.CreateComment(prisma.CommentCreateInput{
		Body: body,
		Author: prisma.UserCreateOneWithoutCommentsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user,
			},
		},
		Post: prisma.PostCreateOneWithoutCommentsInput{
			Connect: &prisma.PostWhereUniqueInput{
				ID: &post,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return newComment, nil
}
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, body string) (*prisma.Comment, error) {
	if id == "" {
		return nil, fmt.Errorf("please provide an id")
	}

	if body == "" {
		return nil, fmt.Errorf("please provide a body")
	}
	updatedComment, err := client.UpdateComment(prisma.CommentUpdateParams{
		Where: prisma.CommentWhereUniqueInput{
			ID: &id,
		},
		Data: prisma.CommentUpdateInput{
			Body: &body,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*prisma.Comment, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePostLike(ctx context.Context, like PostLikeInput) (*prisma.LikePost, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateCommentLike(ctx context.Context, like CommentLikeInput) (*prisma.LikeComment, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateCommentLike(ctx context.Context, id string, quantity int, post string) (*prisma.LikeComment, error) {
	panic("not implemented")
}
func (r *mutationResolver)	UpdatePostLike(ctx context.Context, id string, quantity int, comment string) (*prisma.LikePost, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) Picture(ctx context.Context, obj *prisma.Post) (*graphql.Upload, error) {
	panic("not implemented")
}
func (r *postResolver) CreatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	panic("not implemented")
}
func (r *postResolver) UpdatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	panic("not implemented")
}
func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (*prisma.User, error) {
	panic("not implemented")
}
func (r *postResolver) Comments(ctx context.Context, obj *prisma.Post) ([]prisma.Comment, error) {
	panic("not implemented")
}
func (r *postResolver) Likes(ctx context.Context, obj *prisma.Post) ([]Like, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	users, err := client.Users(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *queryResolver) Posts(ctx context.Context) ([]prisma.Post, error) {
	posts, err := client.Posts(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
func (r *queryResolver) Likes(ctx context.Context) ([]prisma.Like, error) {
	likes, err := client.Likes(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return likes, nil
}
func (r *queryResolver) Comments(ctx context.Context) ([]prisma.Comment, error) {
	comments, err := client.Comments(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return comments, nil
}
func (r *queryResolver) UserByID(ctx context.Context, id string) (*prisma.User, error) {
	user, err := client.User(prisma.UserWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*prisma.User, error) {
	user, err := client.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *queryResolver) Post(ctx context.Context, id string) (*prisma.Post, error) {
	post, err := client.Post(prisma.PostWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return post, nil
}
func (r *queryResolver) CommentLikes(ctx context.Context) ([]prisma.LikeComment, error) {
	panic("not implemented")
}
func (r *queryResolver) PostLikes(ctx context.Context) ([]prisma.LikePost, error) {
	panic("not implemented")
}

func (r *queryResolver) CommentLike(ctx context.Context, id string) (*prisma.LikeComment, error) {
	panic("not implemented")
}
func (r *queryResolver)  PostLike(ctx context.Context, id string) (*prisma.LikePost, error) {
	panic("not implemented")
}

func (r *queryResolver) Comment(ctx context.Context, id string) (*prisma.Comment, error) {
	panic("not implemented")
}
func (r *queryResolver) Me(ctx context.Context) (*prisma.User, error) {
	w := ctx.Value("response").(Auth)
	cookie, err := w.RQ.Cookie("user")
	if err != nil {
		return nil, err
	}

	value := cookie.Value

	currentUser, err := client.User(prisma.UserWhereUniqueInput{
		ID: &value,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return currentUser, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) CreatedAt(ctx context.Context, obj *prisma.User) (*time.Time, error) {
	createdAt := obj.CreatedAt
	t, err := time.Parse(time.RFC3339, createdAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	userPosts, err := client.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Posts(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return userPosts, nil
}
func (r *userResolver) Likes(ctx context.Context, obj *prisma.User) ([]Like, error) {
	panic("not implemented")
	// userLikes, err := client.User(prisma.UserWhereUniqueInput{
	// 	ID: &obj.ID,
	// }).Likes(nil).Exec(ctx)

	// if err != nil {
	// 	return nil, err
	// }

	// return userLikes, nil
}

func (r *userResolver) Comments(ctx context.Context, obj *prisma.User) ([]prisma.Comment, error) {
	userComments, err := client.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Comments(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return userComments, nil
}
