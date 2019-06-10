package invariants

import (
	"math/rand"
	"testing"
)

func TestRandomlyValid100(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := [32]byte{}
		b := [32]byte{}
		rand.Read(b[:])

		RandomlyValid(a[:], b[:], 1)
		if a != b {
			t.Error("100% chance is not working")
			break
		}
	}
}

func TestRandomlyValid0(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := [32]byte{}
		b := [32]byte{}
		rand.Read(b[:])

		RandomlyValid(a[:], b[:], 0)
		if a == b {
			t.Error("0% chance is not working")
			break
		}
	}
}

func TestRandomlyValid01(t *testing.T) {
	count := 1000
	validCases := 0
	for i := 0; i < count; i++ {
		a := [32]byte{}
		b := [32]byte{}
		rand.Read(b[:])

		RandomlyValid(a[:], b[:], 0.01)
		if a == b {
			validCases++
		}
	}
	t.Logf("%d/%d valid cases ~ 1%%", validCases, count)
}
