package fullstack_backend

import (
	"context"
	"fmt"
	// "log"
	// "errors"
	"net/http"
	"time"

	prisma "github.com/german9304/fullstack-backend/prisma-client"
)

var (
	client *prisma.Client = prisma.New(nil)
)

type Resolver struct{}

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

func (r *likeResolver) User(ctx context.Context, obj *prisma.Like) (*prisma.User, error) {
	user, err := client.Like(prisma.LikeWhereUniqueInput{
		ID: &obj.ID,
	}).User().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *likeResolver) Post(ctx context.Context, obj *prisma.Like) (*prisma.Post, error) {
	post, err := client.Like(prisma.LikeWhereUniqueInput{
		ID: &obj.ID,
	}).Post().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return post, nil
}
func (r *likeResolver) CreatedAt(ctx context.Context, obj *prisma.Like) (*time.Time, error) {
	createdAt := obj.CreatedAt
	t, err := time.Parse(time.RFC3339, createdAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (r *likeResolver) UpdatedAt(ctx context.Context, obj *prisma.Like) (*time.Time, error) {
	createdAt := obj.CreatedAt
	t, err := time.Parse(time.RFC3339, createdAt)

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
		return &prisma.User{}, err
	}

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

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, text string) (*prisma.Post, error) {
	updatedPost, err := client.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{
			ID: &id,
		},
		Data: prisma.PostUpdateInput{
			Text: &text,
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

func (r *mutationResolver) CreateLike(ctx context.Context, likeInput LikeInput) (*prisma.Like, error) {
	userID := likeInput.User
	postID := likeInput.Post
	quantity := int32(likeInput.Quantity)

	like, err := client.CreateLike(prisma.LikeCreateInput{
		Quantity: &quantity,
		User: prisma.UserCreateOneWithoutLikesInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &userID,
			},
		},
		Post: prisma.PostCreateOneWithoutLikesInput{
			Connect: &prisma.PostWhereUniqueInput{
				ID: &postID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return like, nil

}

type postResolver struct{ *Resolver }

func (r *postResolver) CreatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	createdAt := obj.CreatedAt
	t, err := time.Parse(time.RFC3339, createdAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (r *postResolver) UpdatedAt(ctx context.Context, obj *prisma.Post) (*time.Time, error) {
	updatedAt := obj.UpdatedAt
	t, err := time.Parse(time.RFC3339, updatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
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
func (r *postResolver) Likes(ctx context.Context, obj *prisma.Post) ([]prisma.Like, error) {
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
func (r *queryResolver) Like(ctx context.Context, id string) (*prisma.Like, error) {
	like, err := client.Like(prisma.LikeWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return like, nil
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
func (r *userResolver) Likes(ctx context.Context, obj *prisma.User) ([]prisma.Like, error) {
	userLikes, err := client.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Likes(nil).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return userLikes, nil
}
