package fullstack_backend

import (
	"context"
	"log"
	"time"

	models "github.com/german9304/fullstack-backend/models"
	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

var (
	client *prisma.Client = prisma.New(nil)
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	postModel *models.PostModel
}

func (r *Resolver) Like() LikeResolver {
	return &likeResolver{r}
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

type likeResolver struct{ *Resolver }

func (r *likeResolver) User(ctx context.Context, obj *prisma.Likes) (*prisma.User, error) {
	panic("not implemented likes")
}
func (r *likeResolver) Post(ctx context.Context, obj *prisma.Likes) (*prisma.Post, error) {
	panic("not implemented likes")
}
func (r *likeResolver) CreatedAt(ctx context.Context, obj *prisma.Likes) (*time.Time, error) {
	panic("not implemented likes")
}
func (r *likeResolver) UpdatedAt(ctx context.Context, obj *prisma.Likes) (*time.Time, error) {
	panic("not implemented likes")
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
		return &prisma.User{}, err
	}

	return user, nil
}
func (r *mutationResolver) Signin(ctx context.Context, email string, password string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) Signout(ctx context.Context) (*Message, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, pstinpt PostInput) (*prisma.Post, error) {
	text := pstinpt.Text
	author := pstinpt.Author

	newPost := client.CreatePost(prisma.PostCreateInput{
		Text: text,
		Author: &prisma.UserCreateOneWithoutPostsInput{
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
func (r *mutationResolver) CreateLike(ctx context.Context, user *string, quantity *int) (*prisma.Likes, error) {
	panic("not implemented likes")
}

type postResolver struct{ *Resolver }

func (r *postResolver) CreatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	log.Println("author here")
	panic("not implemented")
}
func (r *postResolver) UpdatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	log.Println("author here")
	panic("not implemented")
}
func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (*prisma.User, error) {
	postID := obj.ID

	postAuthor, err := client.Post(prisma.PostWhereUniqueInput{
		ID: &postID,
	}).Author().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return postAuthor, nil
}
func (r *postResolver) Likes(ctx context.Context, obj *prisma.Post) ([]prisma.Likes, error) {
	postID := obj.ID

	postLikes, err := client.Post(prisma.PostWhereUniqueInput{
		ID: &postID,
	}).Likes(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return postLikes, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Posts(ctx context.Context) ([]prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Likes(ctx context.Context) ([]prisma.Likes, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id *string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, id *string) (*prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Like(ctx context.Context, id *string) (*prisma.Likes, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) CreatedAt(ctx context.Context, obj *prisma.User) (*time.Time, error) {
	panic("not implemented user")
}
func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	panic("not implemented user")
}
func (r *userResolver) Likes(ctx context.Context, obj *prisma.User) ([]prisma.Likes, error) {
	panic("not implemented user")
}
