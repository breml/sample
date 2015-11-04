package sample

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ReciprocalUint64 is a Sampler, which uses the reciprocal and integer overflow
// The sampling is done by the calculation of the remainder:
//     reciprocal(rate) * probe > reciprocal(rate)
// Reciprocal is calculated with:
//     math.MaxUint64/uint64(rate) + 1
type reciprocalUint64 struct {
	sampleState
	rateuint64 uint64
	reciprocal uint64
}

// NewReciprocalUint64 returns a Modulo Sampler, random number generator is seeded with UnixNano
func NewReciprocalUint64(rate uint64) (Sampler, error) {
	return NewReciprocalUint64Seeded(rate, time.Now().UTC().UnixNano())
}

// NewReciprocalUint64Seeded returns a ReciprocalUint64 Sampler, allow for manual seeding of the random number generator
func NewReciprocalUint64Seeded(rate uint64, seed int64) (Sampler, error) {
	rateuint64 := uint64(rate)
	if rateuint64 == 0 {
		return nil, errors.New("rate must not be 0")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &reciprocalUint64{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}, rateuint64: rateuint64, reciprocal: math.MaxUint64/rateuint64 + 1}, nil
}

func (reciprocal *reciprocalUint64) Sample() bool {
	return reciprocal.SampleFrom(randUint64(reciprocal.rnd))
}

func (reciprocal *reciprocalUint64) SampleFrom(probe uint64) bool {
	probeuint64 := uint64(probe)
	reciprocal.sampleCount++
	if probeuint64 == 0 {
		return true
	}
	if probeuint64 < reciprocal.rateuint64 {
		return false
	}
	if (probeuint64*reciprocal.reciprocal) < reciprocal.reciprocal || reciprocal.rate == 1 {
		reciprocal.trueCount++
		return true
	}
	return false
}

func (reciprocal *reciprocalUint64) String() string {
	type X *reciprocalUint64
	x := X(reciprocal)
	return fmt.Sprintf("%+v", x)
}
