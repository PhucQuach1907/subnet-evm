package transferToken

import (
	"github.com/ava-labs/subnet-evm/eth"
	"github.com/ava-labs/subnet-evm/rpc"
)

// RegisterAPIs registers the token transfer APIs
func RegisterAPIs(backend *eth.EthAPIBackend) []rpc.API {
	api := NewTransferTokenAPI(backend)
	return APIs(api)
} 