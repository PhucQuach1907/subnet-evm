package helloworld

import (
	"context"

	"github.com/ava-labs/subnet-evm/rpc"
)

type PublicAPI struct {}

func (api *PublicAPI) HelloWorld(ctx context.Context) (string, error) {
	return "Hello World", nil
}

// APIs returns the list of APIs this package provides
func APIs(api *PublicAPI) []rpc.API {
	return []rpc.API{
		{
			Namespace: "hello",
			Service:   api,
			Name:      "world",
		},
	}
}