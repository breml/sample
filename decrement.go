package sample

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Decrement is a Sampler, which samples if the counter, initialized by a random number, is decremented to 0
type decrement struct {
	sampleState
	count uint64
}

// NewDecrement returns a Decrement Sampler, random number generator is seeded with UnixNano
func NewDecrement(rate uint64) (Sampler, error) {
	return NewDecrementSeeded(rate, time.Now().UTC().UnixNano())
}

// NewDecrementSeeded returns a Decrement Sampler, allow for manual seeding of the random number generator
func NewDecrementSeeded(rate uint64, seed int64) (Sampler, error) {
	if rate == 0 {
		return nil, errors.New("rate must not be 0")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &decrement{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}, count: decrementSeed(rnd, rate)}, nil
}

func (dec *decrement) Sample() bool {
	dec.sampleCount++
	if dec.count == 0 {
		dec.trueCount++
		dec.count = decrementSeed(dec.rnd, dec.rate)
		return true
	}
	dec.count--
	return false
}

func (dec *decrement) SampleFrom(probe uint64) bool {
	dec.sampleCount++
	if probe%dec.rate == 0 {
		dec.trueCount++
		return true
	}
	return false
}

func (dec *decrement) String() string {
	type X *decrement
	x := X(dec)
	return fmt.Sprintf("%+v", x)
}

func decrementSeed(rnd *rand.Rand, rate uint64) uint64 {
	if rate == 1 {
		return 0
	}
	return randUint64(rnd) % (rate*2 - 2)
}
