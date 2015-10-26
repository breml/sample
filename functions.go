package sample

import (
	"math/rand"
)

func randUint64(rnd *rand.Rand) uint64 {
	return uint64(rnd.Uint32())<<32 + uint64(rnd.Uint32())
}
