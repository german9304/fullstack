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
}

var (
	client *prisma.Client  = prisma.New(nil)
	ctx    context.Context = context.TODO()
	email  string          = "John@mail.com"
	clientGraphql *graphql.Client = graphql.NewClient("http://localhost:8000/")
)

func (fs *FullStackSuite) BeforeTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)
	// name := "John"
	// password := "293902122"
	// age := 32

	// user, _ := client.CreateUser(prisma.UserCreateInput{
	// 	Email:    email,
	// 	Name:     name,
	// 	Password: password,
	// 	Age:      int32(age),
	// }).Exec(ctx)

	// log.Printf("type => %T \n", user)
	// log.Printf("Value => %v \n", user)
}

func (fs *FullStackSuite) AfterTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)
	// user, _ := client.DeleteUser(prisma.UserWhereUniqueInput{
	// 	Email: &email,
	// }).Exec(ctx)
	// log.Printf("Deleted %v \n", user)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (fs *FullStackSuite) TestCreateUser() {
	// make a request
	req := graphql.NewRequest(`
		mutation CreateUserMutation($userinput: UserInput) {
			createUser (usrinpt: $userinput) {
				id
				name
				age
			}
		}
	`)

	usr := UserInput{"mark@mail.com", "Makrk", "2923ij3j3", 32}

	// set any variables
	req.Var("userinput", usr)

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData map[string]interface{}
	if err := clientGraphql.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	log.Printf("result => %v \n", respData)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSetSuite(t *testing.T) {
	suite.Run(t, new(FullStackSuite))
}
