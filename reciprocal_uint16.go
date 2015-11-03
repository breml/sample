//+build ignore

package sample

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ReciprocalUint16 is a Sampler, which uses the reciprocal and integer overflow
// The sampling is done by the calculation of the remainder:
//     reciprocal(rate) * probe > reciprocal(rate)
// Reciprocal is calculated with:
//     math.MaxUint16/uint16(rate) + 1
type reciprocalUint16 struct {
	sampleState
	rateuint16 uint16
	reciprocal uint16
}

// NewReciprocalUint16 returns a Modulo Sampler, random number generator is seeded with UnixNano
func NewReciprocalUint16(rate uint64) (Sampler, error) {
	return NewReciprocalUint16Seeded(rate, time.Now().UTC().UnixNano())
}

// NewReciprocalUint16Seeded returns a ReciprocalUint16 Sampler, allow for manual seeding of the random number generator
func NewReciprocalUint16Seeded(rate uint64, seed int64) (Sampler, error) {
	rateuint16 := uint16(rate)
	if rateuint16 == 0 {
		return nil, errors.New("rate must not be 0")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &reciprocalUint16{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}, rateuint16: rateuint16, reciprocal: math.MaxUint16/rateuint16 + 1}, nil
}

func (reciprocal *reciprocalUint16) Sample() bool {
	return reciprocal.SampleFrom(randUint64(reciprocal.rnd))
}

func (reciprocal *reciprocalUint16) SampleFrom(probe uint64) bool {
	probeuint16 := uint16(probe)
	reciprocal.sampleCount++
	if probeuint16 == 0 {
		return true
	}
	if probeuint16 < reciprocal.rateuint16 {
		return false
	}
	if (probeuint16*reciprocal.reciprocal) < reciprocal.reciprocal || reciprocal.rate == 1 {
		reciprocal.trueCount++
		return true
	}
	return false
}

func (reciprocal *reciprocalUint16) String() string {
	type X *reciprocalUint16
	x := X(reciprocal)
	return fmt.Sprintf("%+v", x)
}
