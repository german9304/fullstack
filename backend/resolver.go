package fullstack_backend

import (
	"context"
	"time"

	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	users []*prisma.User
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

func (r *likeResolver) User(ctx context.Context, obj *prisma.Likes) (string, error) {
	panic("not implemented")
}
func (r *likeResolver) Post(ctx context.Context, obj *prisma.Likes) (string, error) {
	panic("not implemented")
}
func (r *likeResolver) CreatedAt(ctx context.Context, obj *prisma.Likes) (*time.Time, error) {
	panic("not implemented")
}
func (r *likeResolver) UpdatedAt(ctx context.Context, obj *prisma.Likes) (*time.Time, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, usrinpt *UserInput) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, pstinpt *PostInput) (*prisma.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLike(ctx context.Context, user *string, quantity *int) (*prisma.Likes, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) CreatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	panic("not implemented")
}
func (r *postResolver) UpdatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	panic("not implemented")
}
func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (string, error) {
	panic("not implemented")
}
func (r *postResolver) Likes(ctx context.Context, obj *prisma.Post) ([]*prisma.Likes, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*prisma.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Posts(ctx context.Context) ([]*prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Likes(ctx context.Context) ([]*prisma.Likes, error) {
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
	panic("not implemented")
}
func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]*prisma.Post, error) {
	panic("not implemented")
}
func (r *userResolver) Likes(ctx context.Context, obj *prisma.User) ([]*prisma.Likes, error) {
	panic("not implemented")
}
