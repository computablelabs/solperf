package arrayofstructs

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type ctx struct {
	Alloc      core.GenesisAlloc
	Auth       *bind.TransactOpts
	Blockchain *backends.SimulatedBackend
}

type dep struct {
	Address     common.Address
	Contract    *ArrayOfStructs
	Transaction *types.Transaction
}

func Deploy(c *ctx) (*dep, error) {
	// see the generated mappingplusmapping.go classes
	addr, tx, cont, err := DeployArrayOfStructs(c.Auth, c.Blockchain)

	if err != nil {
		return nil, err
	}

	// commit the deploy before returning
	c.Blockchain.Commit()

	return &dep{Address: addr, Contract: cont, Transaction: tx}, nil
}

func setupBlockchain(balance *big.Int) *ctx {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	bc := backends.NewSimulatedBackend(alloc, 4000000) // setting a 4M gas limit here

	return &ctx{
		Alloc:      alloc,
		Auth:       auth,
		Blockchain: bc,
	}
}

// to keep our test main from ballooning in size, abstract the actual setup here
func SetupSession(gasPrice *big.Int, gasLimit uint64, c *ctx, d *dep) *ArrayOfStructsSession {
	return &ArrayOfStructsSession{
		Contract: d.Contract,
		CallOpts: bind.CallOpts{
			From: c.Auth.From,
		},
		TransactOpts: bind.TransactOpts{
			From:     c.Auth.From,
			Signer:   c.Auth.Signer,
			GasLimit: gasLimit,
			GasPrice: gasPrice,
		},
	}
}
