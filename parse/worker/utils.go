package worker

import (
	"strings"

	"github.com/desmos-labs/juno/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// findValidatorByAddr finds a validator by a HEX address given a set of
// Tendermint validators for a particular block. If no validator is found, nil
// is returned.
func findValidatorByAddr(addrHex string, vals *tmctypes.ResultValidators) *tmtypes.Validator {
	for _, val := range vals.Validators {
		if strings.ToLower(addrHex) == strings.ToLower(val.Address.String()) {
			return val
		}
	}

	return nil
}

// sumGasTxs returns the total gas consumed by a set of transactions.
func sumGasTxs(txs []types.Tx) uint64 {
	var totalGas uint64

	for _, tx := range txs {
		totalGas += uint64(tx.GasUsed)
	}

	return totalGas
}
