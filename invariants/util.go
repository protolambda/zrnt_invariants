package invariants

import "encoding/binary"

// make a random byte string valid, but randomly mutate 1 bit, based on given mutation chance.
// chance = 0 results in never valid.
// chance = 1 results in always valid
func RandomlyValid(valid []byte, random []byte, chance float32) {
	chanceRNG := binary.LittleEndian.Uint32(random[:4])
	bit := random[4]
	// make random all valid
	copy(random, valid)
	v := float32(float64(chanceRNG) / float64(^uint32(0)))
	// now mutate random bit based on chance
	if v > chance || chance == 0 {
		random[bit >> 3] ^= 1 << (bit & 0x7)
	}
}

