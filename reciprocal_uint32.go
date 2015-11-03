//+build ignore

package sample

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ReciprocalUint32 is a Sampler, which uses the reciprocal and integer overflow
// The sampling is done by the calculation of the remainder:
//     reciprocal(rate) * probe > reciprocal(rate)
// Reciprocal is calculated with:
//     math.MaxUint32/uint32(rate) + 1
type reciprocalUint32 struct {
	sampleState
	rateuint32 uint32
	reciprocal uint32
}

// NewReciprocalUint32 returns a Modulo Sampler, random number generator is seeded with UnixNano
func NewReciprocalUint32(rate uint64) (Sampler, error) {
	return NewReciprocalUint32Seeded(rate, time.Now().UTC().UnixNano())
}

// NewReciprocalUint32Seeded returns a ReciprocalUint32 Sampler, allow for manual seeding of the random number generator
func NewReciprocalUint32Seeded(rate uint64, seed int64) (Sampler, error) {
	rateuint32 := uint32(rate)
	if rateuint32 == 0 {
		return nil, errors.New("rate must not be 0")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &reciprocalUint32{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}, rateuint32: rateuint32, reciprocal: math.MaxUint32/rateuint32 + 1}, nil
}

func (reciprocal *reciprocalUint32) Sample() bool {
	return reciprocal.SampleFrom(randUint64(reciprocal.rnd))
}

func (reciprocal *reciprocalUint32) SampleFrom(probe uint64) bool {
	probeuint32 := uint32(probe)
	reciprocal.sampleCount++
	if probeuint32 == 0 {
		return true
	}
	if probeuint32 < reciprocal.rateuint32 {
		return false
	}
	if (probeuint32*reciprocal.reciprocal) < reciprocal.reciprocal || reciprocal.rate == 1 {
		reciprocal.trueCount++
		return true
	}
	return false
}

func (reciprocal *reciprocalUint32) String() string {
	type X *reciprocalUint32
	x := X(reciprocal)
	return fmt.Sprintf("%+v", x)
}
