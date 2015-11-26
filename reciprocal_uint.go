//+build generate
//go:generate ./gen_reciprocals.sh 8
//go:generate ./gen_reciprocals.sh 16
//go:generate ./gen_reciprocals.sh 32
//go:generate ./gen_reciprocals.sh 64

package sample

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ReciprocalUint8 is a Sampler, which uses the reciprocal and integer overflow
// The sampling is done by the calculation of the remainder:
//     reciprocal(rate) * probe > reciprocal(rate)
// Reciprocal is calculated with:
//     math.MaxUint8/uint8(rate) + 1
type reciprocalUint8 struct {
	sampleState
	rateuint8  uint8
	reciprocal uint8
}

// NewReciprocalUint8 returns a Modulo Sampler, random number generator is seeded with UnixNano
func NewReciprocalUint8(rate uint64) (Sampler, error) {
	return NewReciprocalUint8Seeded(rate, time.Now().UTC().UnixNano())
}

// NewReciprocalUint8Seeded returns a ReciprocalUint8 Sampler, allow for manual seeding of the random number generator
func NewReciprocalUint8Seeded(rate uint64, seed int64) (Sampler, error) {
	rateuint8 := uint8(rate)
	if rateuint8 == 0 {
		return nil, errors.New("rate must not be 0")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &reciprocalUint8{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}, rateuint8: rateuint8, reciprocal: math.MaxUint8/rateuint8 + 1}, nil
}

func (reciprocal *reciprocalUint8) Sample() bool {
	return reciprocal.SampleFrom(randUint64(reciprocal.rnd))
}

func (reciprocal *reciprocalUint8) SampleFrom(probe uint64) bool {
	probeuint8 := uint8(probe)
	reciprocal.sampleCount++
	if probeuint8 == 0 {
		return true
	}
	if probeuint8 < reciprocal.rateuint8 {
		return false
	}
	if (probeuint8*reciprocal.reciprocal) < reciprocal.reciprocal || reciprocal.rate == 1 {
		reciprocal.trueCount++
		return true
	}
	return false
}

func (reciprocal *reciprocalUint8) String() string {
	type X *reciprocalUint8
	x := X(reciprocal)
	return fmt.Sprintf("%+v", x)
}
