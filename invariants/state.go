package invariants

import (
	"fmt"
	. "github.com/protolambda/zrnt/eth2/beacon"
	. "github.com/protolambda/zrnt/eth2/core"
)

func CheckState(state *BeaconState) error {
	/* Balances and ValidatorRegistry must be the same length */
	if len(state.Balances) != len(state.ValidatorRegistry) {
		return fmt.Errorf("balances/validator-registry length mismatch (%d and %d)", len(state.Balances), len(state.ValidatorRegistry))
	}

	/* Check if the bare minimum of assumed validators is active */
	activeValidators := state.ValidatorRegistry.GetActiveValidatorCount(state.Epoch())
	if activeValidators < uint64(SLOTS_PER_EPOCH*TARGET_COMMITTEE_SIZE) {
		return fmt.Errorf("insufficient active validators %d (out of %d)", activeValidators, len(state.ValidatorRegistry))
	}

	// TODO many more invariants

	return nil
}
