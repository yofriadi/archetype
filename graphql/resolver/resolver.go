package resolver

import (
	"net/http"

	generated "github.com/yofriadi/backend-vanilla/generated/graphql"
	rpcdesign "github.com/yofriadi/backend-vanilla/generated/merchandise/design/v1/designv1connect"
	"github.com/yofriadi/backend-vanilla/graphql/resolver/query"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct{}

// type queryResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver {
	rpcClientDesign := rpcdesign.NewDesignServiceClient(http.DefaultClient, "http://localhost:8080")
	return &query.QueryResolver{
		RPCClientDesign: rpcClientDesign,
	}
}

/* func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutation.MutationResolver{}
} */
