package transferToken

import (
	"context"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/eth"
)

// Service provides token transfer service functionality
type Service struct {
	api *TransferTokenAPI
}

// NewService creates a new Service instance
func NewService(backend *eth.EthAPIBackend) *Service {
	return &Service{
		api: NewTransferTokenAPI(backend),
	}
}

// TransferToken transfers ERC20 tokens
func (s *Service) TransferToken(ctx context.Context, from, to, token common.Address, amount *big.Int, gasPrice *big.Int, privKey string) (common.Hash, error) {
	args := TransferTokenArgs{
		From:     from,
		To:       to,
		Token:    token,
		Amount:   amount,
		GasPrice: gasPrice,
		PrivKey:  privKey,
	}
	return s.api.TransferToken(ctx, args)
} 