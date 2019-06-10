package invariants

import (
	. "github.com/protolambda/zrnt/eth2/beacon"
	"github.com/protolambda/zrnt/eth2/util/ssz"
)

// Idea: pre-process pure random blocks, to get brand new fuzzing corpus items

func PreprocessBlock(pre *BeaconState, block *BeaconBlock) error {

	// 90% chance of hitting a valid block previous-root.
	latestHeaderCopy := pre.LatestBlockHeader
	latestHeaderCopy.StateRoot = ssz.HashTreeRoot(pre, BeaconStateSSZ)
	prevRoot := ssz.SigningRoot(latestHeaderCopy, BeaconBlockHeaderSSZ)
	RandomlyValid(prevRoot[:], block.PreviousBlockRoot[:], 0.9)


	// No BLS yet, but this could be a solution when we want to have real crypto in fuzzing
	//proposerIndex := pre.GetBeaconProposerIndex()
	//validSig := Sign(signing root, getTestingPrivKey(pre.ValidatorRegistry[proposerIndex].Pubkey))
	//RandomlyValid(validSig[:], block.Signature[:], 0.50)

	// make slashing headers valid, sometimes
	// ...

	// etc.

	return nil
}

