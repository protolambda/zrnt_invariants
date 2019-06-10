package invariants

import (
	. "github.com/protolambda/zrnt/eth2/core"
	"testing"
)

func TestGenesisInvariants(t *testing.T) {
	count := uint32(SLOTS_PER_EPOCH * TARGET_COMMITTEE_SIZE)
	// add some margin before having not enough validators.
	count *= 2
	genesis := PrepareGenesisState(count)

	if err := CheckState(genesis); err != nil {
		t.Error(err)
	}
}
