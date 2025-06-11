package transferToken

import (
	"math/big"

	"github.com/ava-labs/libevm/common"
)

// TransferTokenArgs represents the arguments for token transfer
type TransferTokenArgs struct {
	From     common.Address `json:"from"` // Sender address
	To       common.Address `json:"to"` // Receiver address
	Token    common.Address `json:"token"` // ERC20 token contract address
	Amount   *big.Int      `json:"amount"` // Amount of tokens to transfer
	GasPrice *big.Int      `json:"gasPrice,omitempty"` // Gas price in wei
	PrivKey  string        `json:"privKey"` // Private key in hex format
} 

// TransferFromTokenArgs represents the arguments for token transferFrom
type TransferFromTokenArgs struct {
	From     common.Address `json:"from"` // Token owner address
	To       common.Address `json:"to"` // Receiver address
	Token    common.Address `json:"token"` // ERC20 token contract address
	Amount   *big.Int      `json:"amount"` // Amount of tokens to transfer
	GasPrice *big.Int      `json:"gasPrice,omitempty"` // Gas price in wei
	OwnerPrivKey string    `json:"ownerPrivKey"` // Private key of token owner in hex format
} 