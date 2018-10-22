package mappingplusmapping

import (
	"math/big"
	"testing"
)

func TestBecomeOwner(t *testing.T) {
	var STREET_ADDRESS string = "1 Main St. There, CA. 12345"

	t.Log("Entity should be able to become an owner")

	if isOwner, _ := session.IsKooKooRooOwner(context.Auth.From); isOwner != false {
		t.Errorf("Expected Auth.From -> isOwner to be false, got %v", isOwner)
	}

	// now become an owner of 1 location
	becomeOwnerTx, _ := session.BecomeOwner(STREET_ADDRESS)
	context.Blockchain.Commit()
	currentBlock += 1

	receipt, _ := context.Blockchain.TransactionReceipt(nil, becomeOwnerTx.Hash())
	// what did that transaction cost?
	t.Logf("Gas cost to become Owner: %v", becomeOwnerTx.Cost())
	t.Logf("Actual gas used to become Owner: %v", receipt.GasUsed)

	balance, _ := context.Blockchain.BalanceAt(nil, context.Auth.From, big.NewInt(currentBlock))
	t.Logf("Account balance: %v", balance)

	// should be the proud owner of a single location now
	if isOwnerNow, _ := session.IsKooKooRooOwner(context.Auth.From); isOwnerNow != true {
		t.Errorf("Expected Auth.From -> isOwner to be true, got %v", isOwnerNow)
	}

	// assure that it's just one location ATM, NOTE the session method returns bigInt
	if num, _ := session.KooKooRooOwners(context.Auth.From); num.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("Expected Auth.From to own 1 location, got %v", num)
	}

	// we can fetch a location if we want, not too important to the spec, but good to be sure...
	bytes, _ := session.StringToBytes32(STREET_ADDRESS)
	location, _ := session.KooKooRoos(bytes)

	if location.Owner != context.Auth.From {
		t.Errorf("Expected location owner to be Auth.From, got %v", location.Owner)
	}
}

// let's do a small stress test to see what the costs behave like at a small scale
func TestAddingNineteenMore(t *testing.T) {
	t.Log("Entity should be able to become an repeat owner, up to 20")
	// add 19 more and monitor the gas costs and user balance as we go. Main st sure does love chicken...
	LOCATIONS := [19]string{
		"2 Main St. There, CA. 12345",
		"3 Main St. There, CA. 12345",
		"4 Main St. There, CA. 12345",
		"5 Main St. There, CA. 12345",
		"6 Main St. There, CA. 12345",
		"7 Main St. There, CA. 12345",
		"8 Main St. There, CA. 12345",
		"9 Main St. There, CA. 12345",
		"10 Main St. There, CA. 12345",
		"11 Main St. There, CA. 12345",
		"12 Main St. There, CA. 12345",
		"13 Main St. There, CA. 12345",
		"14 Main St. There, CA. 12345",
		"15 Main St. There, CA. 12345",
		"16 Main St. There, CA. 12345",
		"17 Main St. There, CA. 12345",
		"18 Main St. There, CA. 12345",
		"19 Main St. There, CA. 12345",
		"20 Main St. There, CA. 12345",
	}

	for _, v := range LOCATIONS {
		becomeOwnerTx, _ := session.BecomeOwner(v)
		context.Blockchain.Commit()
		currentBlock += 1

		receipt, _ := context.Blockchain.TransactionReceipt(nil, becomeOwnerTx.Hash())
		// what did that transaction cost?
		t.Logf("Gas cost to become Owner of %v: %v", v, becomeOwnerTx.Cost())
		t.Logf("Actual gas used to become Owner of %v: %v", v, receipt.GasUsed)

		balance, _ := context.Blockchain.BalanceAt(nil, context.Auth.From, big.NewInt(currentBlock))
		t.Logf("New account balance: %v", balance)
	}
}
