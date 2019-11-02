package fullstack_backend

import (
	"context"

	"github.com/machinebox/graphql"
)

var (
	// client        *prisma.Client  = prisma.New(nil)
	ctx           context.Context = context.TODO()
	clientGraphql *graphql.Client = graphql.NewClient("http://localhost:8000/")
)