package fullstack_backend

import (
	"context"

	"github.com/machinebox/graphql"
)

var (
	// client        *prisma.Client  = prisma.New(nil)
	ctx           context.Context = context.TODO()
	email         string          = "John@mail.com"
	clientGraphql *graphql.Client = graphql.NewClient("http://localhost:8000/")
)