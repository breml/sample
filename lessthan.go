package sample

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// LessThan is a Sampler, which samples if the probe is lower than boundary, where boundary is calculated as
//     math.MaxUint64 / rate
type lessThan struct {
	sampleState
	boundary uint64
}

// NewLessThan returns a LessThan Sampler, random number generator is seeded with UnixNano
func NewLessThan(rate uint64) (Sampler, error) {
	return NewLessThanSeeded(rate, time.Now().UTC().UnixNano())
}

// NewLessThanSeeded returns a LessThan Sampler, allow for manual seeding of the random number generator
func NewLessThanSeeded(rate uint64, seed int64) (Sampler, error) {
	if rate == 0 {
		return nil, errors.New("rate must not be 0")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &lessThan{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}, boundary: math.MaxUint64 / rate}, nil
}

// Sample returns a sampling decision based on a random number.
func (lt *lessThan) Sample() bool {
	return lt.SampleFrom(randUint64(lt.rnd))
}

// SampleFrom returns a sampling decision based on probe.
func (lt *lessThan) SampleFrom(probe uint64) bool {
	lt.sampleCount++
	if probe < lt.boundary || lt.rate == 1 {
		lt.trueCount++
		return true
	}
	return false
}

// String returns the internal state of the sampler as string
func (lt *lessThan) String() string {
	type X *lessThan
	x := X(lt)
	return fmt.Sprintf("%+v", x)
}
