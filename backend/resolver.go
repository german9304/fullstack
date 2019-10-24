package fullstack_backend

import (
	"context"
	"time"
	// "log"

	"github.com/german9304/fullstack-backend/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	users []*models.User
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

func (r *likeResolver) User(ctx context.Context, obj *models.Like) (string, error) {
	panic("not implemented")
}
func (r *likeResolver) Post(ctx context.Context, obj *models.Like) (string, error) {
	panic("not implemented")
}
func (r *likeResolver) CreatedAt(ctx context.Context, obj *models.Like) (*time.Time, error) {
	panic("not implemented")
}
func (r *likeResolver) UpdatedAt(ctx context.Context, obj *models.Like) (*time.Time, error) {
	panic("not implemented")
}
func (r *likeResolver) Quantity(ctx context.Context, obj *models.Like) (*int, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, usrinpt *UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, pstinpt *PostInput) (*models.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLike(ctx context.Context, user *string, quantity *int) (*models.Like, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	panic("not implemented")
}
func (r *postResolver) Text(ctx context.Context, obj *models.Post) (string, error) {
	panic("not implemented")
}
func (r *postResolver) CreatedAt(ctx context.Context, obj *models.Post) (*time.Time, error) {
	panic("not implemented")
}
func (r *postResolver) UpdatedAt(ctx context.Context, obj *models.Post) (*time.Time, error) {
	panic("not implemented")
}
func (r *postResolver) Author(ctx context.Context, obj *models.Post) (string, error) {
	panic("not implemented")
}
func (r *postResolver) Likes(ctx context.Context, obj *models.Post) ([]*models.Like, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	likes := []*models.Like{}
	posts := []*models.Post{}
	user1 := models.NewUser("2920jki102", "person1@mail.com", "person1", "kwkwkw", 32, posts, likes)
	user2 := models.NewUser("2920jk23ldldld2", "person2@mail.com", "person2", "kw0w0w0w", 35, posts, likes)
	users := []*models.User{user1, user2}
	return users, nil
}
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Likes(ctx context.Context) ([]*models.Like, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id *string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, id *string) (*models.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Like(ctx context.Context, id *string) (*models.Like, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.ID(), nil
}
func (r *userResolver) Email(ctx context.Context, obj *models.User) (string, error) {
	return obj.Email(), nil
}
func (r *userResolver) Name(ctx context.Context, obj *models.User) (string, error) {
	return obj.Name(), nil
}
func (r *userResolver) Password(ctx context.Context, obj *models.User) (string, error) {
	return obj.Password(), nil
}
func (r *userResolver) Age(ctx context.Context, obj *models.User) (int, error) {
	return obj.Age(), nil
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (*time.Time, error) {
	return obj.CreatedAt(), nil
}
func (r *userResolver) Posts(ctx context.Context, obj *models.User) ([]*models.Post, error) {
	return obj.Posts(), nil
}
func (r *userResolver) Likes(ctx context.Context, obj *models.User) ([]*models.Like, error) {
	return obj.Likes(), nil
}
