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

// TransferFromToken transfers ERC20 tokens using transferFrom
func (s *Service) TransferFromToken(ctx context.Context, from, to, token common.Address, amount *big.Int, gasPrice *big.Int, ownerPrivKey string) (common.Hash, error) {
	args := TransferFromTokenArgs{
		From:         from,
		To:           to,
		Token:        token,
		Amount:       amount,
		GasPrice:     gasPrice,
		OwnerPrivKey: ownerPrivKey,
	}
	return s.api.TransferFromToken(ctx, args)
} 