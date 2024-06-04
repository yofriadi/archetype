package query

import (
	rpcClientDesign "github.com/yofriadi/backend-vanilla/generated/merchandise/design/v1/designv1connect"
)

type QueryResolver struct {
	RPCClientDesign rpcClientDesign.DesignServiceClient
}
