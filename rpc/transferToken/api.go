package transferToken

import (
	"context"
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/eth"
	"github.com/ava-labs/subnet-evm/rpc"
)

type TransferTokenAPI struct{
	backend *eth.EthAPIBackend
}

func NewTransferTokenAPI(backend *eth.EthAPIBackend) *TransferTokenAPI {
	return &TransferTokenAPI{backend: backend}
}

// TransferToken transfers ERC20 tokens from one address to another
func (api *TransferTokenAPI) TransferToken(ctx context.Context, args TransferTokenArgs) (common.Hash, error) {
	if args.From == (common.Address{}) {
		return common.Hash{}, errors.New("from address is required")
	}
	if args.To == (common.Address{}) {
		return common.Hash{}, errors.New("to address is required")
	}
	if args.Token == (common.Address{}) {
		return common.Hash{}, errors.New("token address is required")
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return common.Hash{}, errors.New("amount must be greater than 0")
	}
	if args.PrivKey == "" {
		return common.Hash{}, errors.New("private key is required")
	}

	// Set minimum gas price if not provided
	if args.GasPrice == nil {
		args.GasPrice = big.NewInt(25000000000)
	}

	// Calculate token amount with decimals (18 decimals)
	tokenAmount := new(big.Int).Mul(args.Amount, big.NewInt(1000000000000000000))

	// Function signature: transfer(address,uint256)
	data := []byte{0xa9, 0x05, 0x9c, 0xbb}
	data = append(data, common.LeftPadBytes(args.To.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(tokenAmount.Bytes(), 32)...)

	// Get nonce from backend
	nonce, err := api.backend.GetPoolNonce(ctx, args.From)
	if err != nil {
		return common.Hash{}, err
	}

	// Get chain ID
	chainConfig := api.backend.ChainConfig()
	chainID := chainConfig.ChainID

	// Create new transaction with correct nonce
	tx := types.NewTransaction(
		nonce,                  // Nonce from backend
		args.Token,             // To
		big.NewInt(0),          // Value (0 vì là ERC20 transfer)
		100000,                 // Gas limit
		args.GasPrice,          // Gas price
		data,                   // Data (transfer function)
	)

	// Convert private key to crypto.PrivateKey
	privKey, err := crypto.HexToECDSA(args.PrivKey)
	if err != nil {
		return common.Hash{}, err
	}

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return common.Hash{}, err
	}

	// Send raw transaction
	err = api.backend.SendTx(ctx, signedTx)
	if err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), nil
}

// TransferFromToken transfers ERC20 tokens from one address to another using transferFrom
func (api *TransferTokenAPI) TransferFromToken(ctx context.Context, args TransferFromTokenArgs) (common.Hash, error) {
	if args.From == (common.Address{}) {
		return common.Hash{}, errors.New("from address is required")
	}
	if args.To == (common.Address{}) {
		return common.Hash{}, errors.New("to address is required")
	}
	if args.Token == (common.Address{}) {
		return common.Hash{}, errors.New("token address is required")
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return common.Hash{}, errors.New("amount must be greater than 0")
	}
	if args.OwnerPrivKey == "" {
		return common.Hash{}, errors.New("private key is required")
	}

	// Set minimum gas price if not provided
	if args.GasPrice == nil {
		args.GasPrice = big.NewInt(25000000000)
	}

	// Calculate token amount with decimals (18 decimals)
	tokenAmount := new(big.Int).Mul(args.Amount, big.NewInt(1000000000000000000))

	// Function signature: transferFrom(address,address,uint256)
	data := []byte{0x23, 0xb8, 0x72, 0xdd}
	data = append(data, common.LeftPadBytes(args.From.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(args.To.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(tokenAmount.Bytes(), 32)...)

	// Get spender address from private key
	privKey, err := crypto.HexToECDSA(args.OwnerPrivKey)
	if err != nil {
		return common.Hash{}, err
	}
	spenderAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	// Get nonce from backend
	nonce, err := api.backend.GetPoolNonce(ctx, spenderAddr)
	if err != nil {
		return common.Hash{}, err
	}

	// Get chain ID
	chainConfig := api.backend.ChainConfig()
	chainID := chainConfig.ChainID

	// Create new transaction with correct nonce
	tx := types.NewTransaction(
		nonce,                  // Nonce from backend
		args.Token,             // To
		big.NewInt(0),          // Value (0 vì là ERC20 transfer)
		100000,                 // Gas limit
		args.GasPrice,          // Gas price
		data,                   // Data (transferFrom function)
	)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return common.Hash{}, err
	}

	// Send raw transaction
	err = api.backend.SendTx(ctx, signedTx)
	if err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), nil
}

// APIs returns the list of APIs this package provides
func APIs(api *TransferTokenAPI) []rpc.API {
	return []rpc.API{
		{
			Namespace: "custom",
			Service:   api,
			Name:      "transferToken",
		},
	}
}