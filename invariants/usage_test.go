package invariants

import (
	"bytes"
	"fmt"
	. "github.com/protolambda/zrnt/eth2/beacon"
	"github.com/protolambda/zrnt/eth2/beacon/transition"
	"github.com/protolambda/zssz"
	"math/rand"
	"testing"
)

func TestTransitions(t *testing.T) {
	lastValid := PrepareGenesisState(1000)

	blockSSZ := BeaconBlockSSZ
	minFuzzInputLen := blockSSZ.FuzzReqLen()
	data := make([]byte, minFuzzInputLen * 10)


	for i := 0; i < 10; i++ {
		pre := lastValid.Copy()
		fmt.Printf("creating random block %d and applying it to pre-state (slot %d)\n", i, pre.Slot)

		block := &BeaconBlock{}
		fuzzLen := minFuzzInputLen + uint32(rand.Intn(len(data) - int(minFuzzInputLen)))
		// get some random bytes
		rand.Read(data[:fuzzLen])

		// use the random bytes to create a fully randomized block instance with
		if err, _ := zssz.DecodeFuzzBytes(bytes.NewReader(data), fuzzLen, block, blockSSZ); err != nil {
			t.Fatal(fmt.Errorf("could not create random block, %v\n", err))
		}

		// pre-process the block, we want it to have valid crypto properties some amount of the time
		if err := PreprocessBlock(pre, block); err != nil {
			t.Fatal(fmt.Errorf("could not preprocess block: %v\n", err))
		}

		// For local testing purposes, limit the slot range, no timeout / memory restrictions here...
		block.Slot = block.Slot % 200

		// try to transition using the block
		err := transition.StateTransition(pre, block, false) // don't verify state root of block
		if err != nil {
			fmt.Printf("tried processing random block %d, but was invalid transition: %v\n", i, err)
		} else {
			// take post state (processed pre) as new starting state.
			lastValid = pre.Copy()
		}
	}
}
