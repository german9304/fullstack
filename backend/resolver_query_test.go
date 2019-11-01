package fullstack_backend

import (
	"log"
	"testing"

	// prisma "github.com/german9304/fullstack-backend/prisma-client"
	// "github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

type FullstackSuiteQuery struct {
	suite.Suite
}

func (fs *FullstackSuiteQuery) SetupTest() {
	log.Println("Set up test runs once")
}

func (fs *FullstackSuiteQuery) BeforeTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)

}

func (fs *FullstackSuiteQuery) AfterTest(suiteName, testName string) {
	log.Printf("s: %v, t: %v \n", suiteName, testName)
}

func (fs *FullstackSuiteQuery) TestQueryUsers() {
	fs.Assert().Equal(5, 5)
}

func (fs *FullstackSuiteQuery) TestQueryPosts() {
	fs.Assert().Equal(5, 5)
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(FullstackSuiteQuery))
}
