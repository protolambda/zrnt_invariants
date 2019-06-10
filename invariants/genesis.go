package invariants

import (
	. "github.com/protolambda/zrnt/eth2/beacon"
	"github.com/protolambda/zrnt/eth2/beacon/genesis"
	. "github.com/protolambda/zrnt/eth2/core"
	. "github.com/protolambda/zrnt/eth2/util/hashing"
	"github.com/protolambda/zrnt/eth2/util/math"
	"github.com/protolambda/zrnt/eth2/util/merkle"
	"github.com/protolambda/zrnt/eth2/util/ssz"
	"math/rand"
)

func PrepareGenesisState(validatorCount uint32) *BeaconState {
	genesisTime := Timestamp(1000000000)
	// RNG used to create private keys for test state
	rng := rand.New(rand.NewSource(0xDEADBEEF))

	privKeys := make([][32]byte, 0, validatorCount)
	deposits := make([]Deposit, 0, validatorCount)
	depRoots := make([]Root, 0, validatorCount)
	for i := uint32(0); i < validatorCount; i++ {
		// create a random 32 byte private key. We're not using real crypto yet.
		privKey := [32]byte{}
		rng.Read(privKey[:])
		privKeys = append(privKeys, privKey)
		// simply derive pubkey and withdraw creds, not real thing yet
		pubKey := BLSPubkey{}
		h := Hash(privKey[:])
		// make first 16 bytes recognizable for debugging.
		copy(pubKey[:16], privKey[:16])
		copy(pubKey[16:], h[:])

		withdrawCreds := Hash(append(h[:], 1))
		dep := Deposit{
			Proof: [DEPOSIT_CONTRACT_TREE_DEPTH]Root{},
			Index: DepositIndex(i),
			Data: DepositData{
				Pubkey:                pubKey,
				WithdrawalCredentials: withdrawCreds,
				Amount:                MAX_EFFECTIVE_BALANCE,
				Signature:             BLSSignature{1, 2, 3}, // BLS not yet implemented
			},
		}
		depLeafHash := ssz.HashTreeRoot(&dep.Data, DepositDataSSZ)
		deposits = append(deposits, dep)
		depRoots = append(depRoots, depLeafHash)
	}
	for i := 0; i < len(deposits); i++ {
		copy(deposits[i].Proof[:], merkle.ConstructProof(depRoots, uint64(i), uint8(DEPOSIT_CONTRACT_TREE_DEPTH)))
	}
	power2 := math.NextPowerOfTwo(uint64(len(depRoots)))
	depositsRoot := merkle.MerkleRoot(depRoots)
	// Now pad with zero branches to complete depth.
	buf := [64]byte{}
	for i := power2; i < (1 << DEPOSIT_CONTRACT_TREE_DEPTH); i <<= 1 {
		copy(buf[0:32], depositsRoot[:])
		depositsRoot = Hash(buf[:])
	}

	eth1Data := Eth1Data{
		DepositRoot: depositsRoot,
		BlockHash:   Root{42}, // TODO eth1 simulation
	}
	return genesis.GetGenesisBeaconState(deposits, genesisTime, eth1Data)
}
