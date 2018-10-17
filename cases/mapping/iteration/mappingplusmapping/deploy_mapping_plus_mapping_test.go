package mappingplusmapping

import (
	"math/big"
	"os"
	"testing"
)

// vars declared here have package scope, like our session
var session *MappingPlusMappingSession

// in order to setup a session, we still need a few things to be done
var context *ctx
var deployed *dep
var deployedError error

func TestDeployMappingPlusMapping(t *testing.T) {
	t.Log("MappingPlusMapping contract should deploy correctly")

	if deployedError != nil {
		t.Fatalf("Failed to deploy contract: %v", deployedError)
	}

	if len(deployed.Address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address, but got an empty byte array")
	}

	// let's examine the deployment transaction's cost
	t.Logf("Cost of the deployment TX: %v", deployed.Transaction.Cost())
}

// remember, test main is run once per package so its a good place to init package vars
func TestMain(m *testing.M) {
	// see ./helpers
	context = setupBlockchain(big.NewInt(1000000000000000000)) // 1 ETH in wei
	deployed, deployedError = Deploy(context)
	// arg one: eth gas station says going gwei rate for gas price is 1.5 Gwei, we'll round up to 2
	// setting a 1M Wei gas limit. NOTE GasLimit is expected to be a uint64
	session = SetupSession(big.NewInt(2), 1000000, context, deployed)

	code := m.Run()
	os.Exit(code)
}
