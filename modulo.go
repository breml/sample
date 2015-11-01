package sample

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Modulo is a Sampler, which uses the modulo operation
// The sampling is done by the calculation of the remainder:
//     probe % rate == 0
type modulo struct {
	sampleState
}

// NewModulo returns a Modulo Sampler, random number generator is seeded with UnixNano
func NewModulo(rate uint64) (Sampler, error) {
	return NewModuloSeeded(rate, time.Now().UTC().UnixNano())
}

// NewModuloSeeded returns a Modulo Sampler, allow for manual seeding of the random number generator
func NewModuloSeeded(rate uint64, seed int64) (Sampler, error) {
	if rate == 0 {
		return nil, errors.New("rate must not be 0")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &modulo{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}}, nil
}

func (mod *modulo) Sample() bool {
	return mod.SampleFrom(randUint64(mod.rnd))
}

func (mod *modulo) SampleFrom(probe uint64) bool {
	mod.sampleCount++
	if probe%mod.rate == 0 {
		mod.trueCount++
		return true
	}
	return false
}

func (mod *modulo) String() string {
	type X *modulo
	x := X(mod)
	return fmt.Sprintf("%+v", x)
}
